package httpapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"appfactory/service-manager/internal/runtime"
	sharedhealth "appfactory/shared-go/health"
	"appfactory/shared-go/httpx"
)

type ServiceRuntime struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"`
	Profile string `json:"profile"`
}

type HealthResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Status  string `json:"status"`
	Profile string `json:"profile"`
}

type operationLocker struct {
	mu    sync.Mutex
	locks map[string]bool
}

type releaseOperationStatus struct {
	Operation       string    `json:"operation"`
	ProductSlug     string    `json:"product_slug"`
	TargetType      string    `json:"target_type"`
	TargetVersionID string    `json:"target_version_id,omitempty"`
	Environment     string    `json:"environment,omitempty"`
	Operator        string    `json:"operator,omitempty"`
	Status          string    `json:"status"`
	StartedAt       time.Time `json:"started_at"`
	CompletedAt     time.Time `json:"completed_at,omitempty"`
	Error           string    `json:"error,omitempty"`
}

type operationTracker struct {
	mu      sync.Mutex
	current map[string]releaseOperationStatus
	history []releaseOperationStatus
	path    string
}

type releasePromoteRequest struct {
	ProductSlug     string `json:"product_slug"`
	TargetType      string `json:"target_type"`
	TargetVersionID string `json:"target_version_id"`
	Environment     string `json:"environment"`
	Operator        string `json:"operator"`
}

func NewRouter() http.Handler {
	manager, err := runtime.NewManagerFromConfig("configs/local.yaml")
	if err != nil {
		manager = runtime.NewManager([]runtime.ConfigService{
			{Name: "account-service", Command: "./bin/account-service", WorkDir: "../account-service", Address: "http://localhost:8081"},
			{Name: "upgrade-service", Command: "./bin/upgrade-service", WorkDir: "../upgrade-service", Address: "http://localhost:8082"},
		}, "local")
	}
	return NewRouterWithManager(manager)
}

func NewRouterWithManager(manager *runtime.Manager) http.Handler {
	tracker, err := newOperationTracker(filepath.Join("data", "service-manager-operations.json"))
	if err != nil {
		tracker = &operationTracker{current: map[string]releaseOperationStatus{}}
	}
	return NewRouterWithManagerAndTracker(manager, tracker)
}

func NewRouterWithManagerAndTracker(manager *runtime.Manager, tracker *operationTracker) http.Handler {
	mux := http.NewServeMux()
	httpClient := &http.Client{Timeout: 2 * time.Second}
	locker := &operationLocker{locks: map[string]bool{}}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, sharedhealth.Snapshot{
			Service: "service-manager",
			Status:  "ok",
			Checks: map[string]string{
				"http": "ok",
			},
		})
	})

	mux.HandleFunc("/v1/services", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"services": toServiceRuntimes(manager.List())})
	})
	mux.HandleFunc("/v1/services/start", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		service, err := manager.Start(req.Name)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusAccepted, service)
	})
	mux.HandleFunc("/v1/services/stop", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		service, err := manager.Stop(req.Name)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusAccepted, service)
	})
	mux.HandleFunc("/v1/services/restart", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		service, err := manager.Restart(req.Name)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusAccepted, service)
	})
	mux.HandleFunc("/v1/services/health", func(w http.ResponseWriter, r *http.Request) {
		services := toServiceRuntimes(manager.List())
		results := make([]HealthResult, 0, len(services))
		for _, service := range services {
			status := service.Status
			resp, err := httpClient.Get(service.Address + "/healthz")
			if err == nil && resp.StatusCode == http.StatusOK {
				status = "ok"
				resp.Body.Close()
			} else if err != nil {
				status = "unreachable"
			}
			results = append(results, HealthResult{
				Name:    service.Name,
				Address: service.Address,
				Status:  status,
				Profile: service.Profile,
			})
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"services": results})
	})
	mux.HandleFunc("/v1/services/switch-profile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req struct {
			Profile string `json:"profile"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		manager.SwitchProfile(req.Profile)
		httpx.WriteJSON(w, http.StatusAccepted, map[string]string{
			"message": "switch-profile accepted",
			"profile": req.Profile,
		})
	})
	mux.HandleFunc("/v1/services/targets", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"targets": toServiceRuntimes(manager.List())})
	})
	mux.HandleFunc("/v1/releases/targets", func(w http.ResponseWriter, r *http.Request) {
		proxyUpgradeRequest(w, r, manager, httpClient, http.MethodGet, "/v1/targets/active")
	})
	mux.HandleFunc("/v1/releases/operations/current", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"operations": tracker.Current()})
	})
	mux.HandleFunc("/v1/releases/operations/history", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"operations": tracker.History()})
	})
	mux.HandleFunc("/v1/releases/history", func(w http.ResponseWriter, r *http.Request) {
		proxyUpgradeRequest(w, r, manager, httpClient, http.MethodGet, "/v1/releases")
	})
	mux.HandleFunc("/v1/releases/create", func(w http.ResponseWriter, r *http.Request) {
		proxyLockedUpgradeRequest(w, r, manager, httpClient, locker, http.MethodPost, "/v1/releases", func(payload []byte) (string, error) {
			var req struct {
				ProductSlug string `json:"product_slug"`
				TargetType  string `json:"target_type"`
			}
			if err := json.Unmarshal(payload, &req); err != nil {
				return "", err
			}
			return releaseLockKey(req.ProductSlug, req.TargetType), nil
		}, tracker, func(payload []byte) releaseOperationStatus {
			var req struct {
				ProductSlug  string `json:"product_slug"`
				TargetType   string `json:"target_type"`
				VersionLabel string `json:"version_label"`
			}
			_ = json.Unmarshal(payload, &req)
			return releaseOperationStatus{
				Operation:   "create-release",
				ProductSlug: req.ProductSlug,
				TargetType:  req.TargetType,
				Status:      "running",
			}
		})
	})
	mux.HandleFunc("/v1/deployments/history", func(w http.ResponseWriter, r *http.Request) {
		proxyUpgradeRequest(w, r, manager, httpClient, http.MethodGet, "/v1/deployments")
	})
	mux.HandleFunc("/v1/deployments/create", func(w http.ResponseWriter, r *http.Request) {
		proxyLockedUpgradeRequest(w, r, manager, httpClient, locker, http.MethodPost, "/v1/deployments", func(payload []byte) (string, error) {
			var req struct {
				TargetVersionID string `json:"target_version_id"`
			}
			if err := json.Unmarshal(payload, &req); err != nil {
				return "", err
			}
			return "deployment:" + req.TargetVersionID, nil
		}, tracker, func(payload []byte) releaseOperationStatus {
			var req struct {
				TargetVersionID string `json:"target_version_id"`
				Environment     string `json:"environment"`
			}
			_ = json.Unmarshal(payload, &req)
			return releaseOperationStatus{
				Operation:       "create-deployment",
				TargetVersionID: req.TargetVersionID,
				Environment:     req.Environment,
				Status:          "running",
			}
		})
	})
	mux.HandleFunc("/v1/releases/switches/history", func(w http.ResponseWriter, r *http.Request) {
		proxyUpgradeRequest(w, r, manager, httpClient, http.MethodGet, "/v1/switches")
	})
	mux.HandleFunc("/v1/releases/rollbacks/history", func(w http.ResponseWriter, r *http.Request) {
		proxyUpgradeRequest(w, r, manager, httpClient, http.MethodGet, "/v1/rollbacks")
	})
	mux.HandleFunc("/v1/releases/switch", func(w http.ResponseWriter, r *http.Request) {
		proxyLockedUpgradeRequest(w, r, manager, httpClient, locker, http.MethodPost, "/v1/switches", func(payload []byte) (string, error) {
			var req struct {
				ProductSlug string `json:"product_slug"`
				TargetType  string `json:"target_type"`
			}
			if err := json.Unmarshal(payload, &req); err != nil {
				return "", err
			}
			return releaseLockKey(req.ProductSlug, req.TargetType), nil
		}, tracker, func(payload []byte) releaseOperationStatus {
			var req struct {
				ProductSlug     string `json:"product_slug"`
				TargetType      string `json:"target_type"`
				TargetVersionID string `json:"to_version_id"`
				Operator        string `json:"operator"`
			}
			_ = json.Unmarshal(payload, &req)
			return releaseOperationStatus{
				Operation:       "switch",
				ProductSlug:     req.ProductSlug,
				TargetType:      req.TargetType,
				TargetVersionID: req.TargetVersionID,
				Operator:        req.Operator,
				Status:          "running",
			}
		})
	})
	mux.HandleFunc("/v1/releases/rollback", func(w http.ResponseWriter, r *http.Request) {
		proxyLockedUpgradeRequest(w, r, manager, httpClient, locker, http.MethodPost, "/v1/rollbacks", func(payload []byte) (string, error) {
			var req struct {
				ProductSlug string `json:"product_slug"`
				TargetType  string `json:"target_type"`
			}
			if err := json.Unmarshal(payload, &req); err != nil {
				return "", err
			}
			return releaseLockKey(req.ProductSlug, req.TargetType), nil
		}, tracker, func(payload []byte) releaseOperationStatus {
			var req struct {
				ProductSlug     string `json:"product_slug"`
				TargetType      string `json:"target_type"`
				TargetVersionID string `json:"rolled_back_to_version_id"`
				Operator        string `json:"operator"`
			}
			_ = json.Unmarshal(payload, &req)
			return releaseOperationStatus{
				Operation:       "rollback",
				ProductSlug:     req.ProductSlug,
				TargetType:      req.TargetType,
				TargetVersionID: req.TargetVersionID,
				Operator:        req.Operator,
				Status:          "running",
			}
		})
	})
	mux.HandleFunc("/v1/releases/promote", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		var req releasePromoteRequest
		if err := json.Unmarshal(payload, &req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		lockKey := releaseLockKey(req.ProductSlug, req.TargetType)
		if !locker.Acquire(lockKey) {
			httpx.WriteJSON(w, http.StatusConflict, map[string]string{"error": "release operation already in progress"})
			return
		}
		defer locker.Release(lockKey)
		status := releaseOperationStatus{
			Operation:       "promote",
			ProductSlug:     req.ProductSlug,
			TargetType:      req.TargetType,
			TargetVersionID: req.TargetVersionID,
			Environment:     req.Environment,
			Operator:        req.Operator,
			Status:          "running",
			StartedAt:       time.Now().UTC(),
		}
		tracker.Start(lockKey, status)

		deploymentPayload, _ := json.Marshal(map[string]any{
			"target_version_id": req.TargetVersionID,
			"environment":       req.Environment,
			"status":            "deployed",
		})
		if _, _, err := performUpgradeRequest(r, manager, httpClient, http.MethodPost, "/v1/deployments", deploymentPayload); err != nil {
			status.Status = "failed"
			status.CompletedAt = time.Now().UTC()
			status.Error = err.Error()
			tracker.Finish(lockKey, status)
			httpx.WriteJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
			return
		}
		switchPayload, _ := json.Marshal(map[string]any{
			"product_slug":  req.ProductSlug,
			"target_type":   req.TargetType,
			"to_version_id": req.TargetVersionID,
			"operator":      req.Operator,
		})
		httpStatus, responseBody, err := performUpgradeRequest(r, manager, httpClient, http.MethodPost, "/v1/switches", switchPayload)
		if err != nil {
			failedStatus := releaseOperationStatus{
				Operation:       "promote",
				ProductSlug:     req.ProductSlug,
				TargetType:      req.TargetType,
				TargetVersionID: req.TargetVersionID,
				Environment:     req.Environment,
				Operator:        req.Operator,
				Status:          "failed",
				StartedAt:       time.Now().UTC(),
				CompletedAt:     time.Now().UTC(),
				Error:           err.Error(),
			}
			tracker.Finish(lockKey, failedStatus)
			httpx.WriteJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
			return
		}
		completedStatus := releaseOperationStatus{
			Operation:       "promote",
			ProductSlug:     req.ProductSlug,
			TargetType:      req.TargetType,
			TargetVersionID: req.TargetVersionID,
			Environment:     req.Environment,
			Operator:        req.Operator,
			Status:          "completed",
			StartedAt:       time.Now().UTC(),
			CompletedAt:     time.Now().UTC(),
		}
		tracker.Finish(lockKey, completedStatus)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		_, _ = w.Write(responseBody)
	})

	return mux
}

func toServiceRuntimes(services []runtime.ManagedService) []ServiceRuntime {
	results := make([]ServiceRuntime, 0, len(services))
	for _, service := range services {
		results = append(results, ServiceRuntime{
			Name:    service.Name,
			Address: service.Address,
			Status:  service.Status,
			Profile: service.Profile,
		})
	}
	return results
}

func proxyUpgradeRequest(
	w http.ResponseWriter,
	r *http.Request,
	manager *runtime.Manager,
	httpClient *http.Client,
	expectedMethod string,
	path string,
) {
	if r.Method != expectedMethod {
		httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	upgradeService, ok := findService(manager.List(), "upgrade-service")
	if !ok {
		httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "upgrade-service not configured"})
		return
	}

	var body io.Reader = http.NoBody
	if r.Body != nil && expectedMethod != http.MethodGet {
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		body = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(r.Context(), expectedMethod, upgradeService.Address+path, body)
	if err != nil {
		httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		httpx.WriteJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		httpx.WriteJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	_, _ = w.Write(responseBody)
}

func proxyLockedUpgradeRequest(
	w http.ResponseWriter,
	r *http.Request,
	manager *runtime.Manager,
	httpClient *http.Client,
	locker *operationLocker,
	expectedMethod string,
	path string,
	lockKeyFn func([]byte) (string, error),
	tracker *operationTracker,
	statusFn func([]byte) releaseOperationStatus,
) {
	if r.Method != expectedMethod {
		httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	lockKey, err := lockKeyFn(payload)
	if err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	if !locker.Acquire(lockKey) {
		httpx.WriteJSON(w, http.StatusConflict, map[string]string{"error": "release operation already in progress"})
		return
	}
	defer locker.Release(lockKey)
	status := statusFn(payload)
	status.StartedAt = time.Now().UTC()
	tracker.Start(lockKey, status)
	defer func() {
		tracker.Finish(lockKey, status)
	}()

	httpStatus, responseBody, err := performUpgradeRequest(r, manager, httpClient, expectedMethod, path, payload)
	if err != nil {
		status.Status = "failed"
		status.Error = err.Error()
		httpx.WriteJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		return
	}
	status.Status = "completed"
	status.CompletedAt = time.Now().UTC()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(responseBody)
}

func performUpgradeRequest(
	r *http.Request,
	manager *runtime.Manager,
	httpClient *http.Client,
	method string,
	path string,
	payload []byte,
) (int, []byte, error) {
	upgradeService, ok := findService(manager.List(), "upgrade-service")
	if !ok {
		return 0, nil, io.EOF
	}

	var body io.Reader = http.NoBody
	if method != http.MethodGet {
		body = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(r.Context(), method, upgradeService.Address+path, body)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, responseBody, nil
}

func (l *operationLocker) Acquire(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.locks[key] {
		return false
	}
	l.locks[key] = true
	return true
}

func (l *operationLocker) Release(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.locks, key)
}

func releaseLockKey(productSlug, targetType string) string {
	return productSlug + ":" + targetType
}

type operationTrackerState struct {
	Current map[string]releaseOperationStatus `json:"current"`
	History []releaseOperationStatus          `json:"history"`
}

func newOperationTracker(path string) (*operationTracker, error) {
	tracker := &operationTracker{
		current: map[string]releaseOperationStatus{},
		history: []releaseOperationStatus{},
		path:    path,
	}
	if path == "" {
		return tracker, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return tracker, nil
		}
		return nil, err
	}
	var state operationTrackerState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	if state.Current != nil {
		tracker.current = state.Current
	}
	if state.History != nil {
		tracker.history = state.History
	}
	return tracker, nil
}

func (t *operationTracker) Start(key string, status releaseOperationStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.current[key] = status
	t.persistLocked()
}

func (t *operationTracker) Finish(key string, status releaseOperationStatus) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if status.CompletedAt.IsZero() {
		status.CompletedAt = time.Now().UTC()
	}
	delete(t.current, key)
	t.history = append([]releaseOperationStatus{status}, t.history...)
	if len(t.history) > 20 {
		t.history = t.history[:20]
	}
	t.persistLocked()
}

func (t *operationTracker) Current() []releaseOperationStatus {
	t.mu.Lock()
	defer t.mu.Unlock()
	ops := make([]releaseOperationStatus, 0, len(t.current))
	for _, status := range t.current {
		ops = append(ops, status)
	}
	return ops
}

func (t *operationTracker) History() []releaseOperationStatus {
	t.mu.Lock()
	defer t.mu.Unlock()
	ops := make([]releaseOperationStatus, len(t.history))
	copy(ops, t.history)
	return ops
}

func (t *operationTracker) persistLocked() {
	if t.path == "" {
		return
	}
	state := operationTrackerState{
		Current: t.current,
		History: t.history,
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return
	}
	if err := os.MkdirAll(filepath.Dir(t.path), 0o755); err != nil {
		return
	}
	_ = os.WriteFile(t.path, data, 0o644)
}

func findService(services []runtime.ManagedService, name string) (runtime.ManagedService, bool) {
	for _, service := range services {
		if service.Name == name {
			return service, true
		}
	}
	return runtime.ManagedService{}, false
}
