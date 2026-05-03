package httpapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
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
	mux := http.NewServeMux()
	httpClient := &http.Client{Timeout: 2 * time.Second}

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
	mux.HandleFunc("/v1/releases/switch", func(w http.ResponseWriter, r *http.Request) {
		proxyUpgradeRequest(w, r, manager, httpClient, http.MethodPost, "/v1/switches")
	})
	mux.HandleFunc("/v1/releases/rollback", func(w http.ResponseWriter, r *http.Request) {
		proxyUpgradeRequest(w, r, manager, httpClient, http.MethodPost, "/v1/rollbacks")
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

func findService(services []runtime.ManagedService, name string) (runtime.ManagedService, bool) {
	for _, service := range services {
		if service.Name == name {
			return service, true
		}
	}
	return runtime.ManagedService{}, false
}
