package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"appfactory/service-manager/internal/runtime"
)

func TestServicesEndpointReturnsRegisteredServices(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/services", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "\"account-service\"") || !strings.Contains(body, "\"upgrade-service\"") {
		t.Fatalf("expected registered services in response body, got %s", body)
	}
}

func TestHealthEndpointAggregatesReachableServices(t *testing.T) {
	healthyService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"service":"account-service","status":"ok"}`))
	}))
	defer healthyService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "account-service", Command: "sleep 30", WorkDir: ".", Address: healthyService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	req := httptest.NewRequest(http.MethodGet, "/v1/services/health", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "\"account-service\"") || !strings.Contains(body, "\"ok\"") {
		t.Fatalf("expected aggregated service health in response body, got %s", body)
	}
}

func TestReleaseSwitchProxyUsesUpgradeService(t *testing.T) {
	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/switches" {
			t.Fatalf("expected /v1/switches, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		_, _ = w.Write([]byte(`{"target_type":"client","latest_version":"26.2.20.06"}`))
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/releases/switch",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","to_version_id":"rv-1","operator":"service-manager"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK && rec.Code != http.StatusCreated {
		t.Fatalf("expected proxied success status, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "\"latest_version\":\"26.2.20.06\"") {
		t.Fatalf("expected proxied release switch response, got %s", rec.Body.String())
	}
}

func TestReleaseTargetsProxyUsesUpgradeService(t *testing.T) {
	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/targets/active" {
			t.Fatalf("expected /v1/targets/active, got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET, got %s", r.Method)
		}
		_, _ = w.Write([]byte(`{"client":{"latest_version":"26.2.20.06"},"service":{"latest_version":"26.2.20.03"}}`))
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	req := httptest.NewRequest(http.MethodGet, "/v1/releases/targets", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected proxied success status, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "\"latest_version\":\"26.2.20.06\"") {
		t.Fatalf("expected proxied active target response, got %s", rec.Body.String())
	}
}

func TestReleaseHistoryProxyUsesUpgradeService(t *testing.T) {
	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/releases" {
			t.Fatalf("expected /v1/releases, got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET, got %s", r.Method)
		}
		_, _ = w.Write([]byte(`{"releases":[{"version_label":"26.2.20.06"}]}`))
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	req := httptest.NewRequest(http.MethodGet, "/v1/releases/history", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected proxied success status, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "\"version_label\":\"26.2.20.06\"") {
		t.Fatalf("expected proxied release history response, got %s", rec.Body.String())
	}
}

func TestCreateReleaseProxyUsesUpgradeService(t *testing.T) {
	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/releases" {
			t.Fatalf("expected /v1/releases, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		_, _ = w.Write([]byte(`{"id":"release-1","version_label":"26.2.20.08"}`))
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/releases/create",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","version_label":"26.2.20.08","build_number":8,"upgrade_url":"https://example.com/client/26.2.20.08"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK && rec.Code != http.StatusCreated {
		t.Fatalf("expected proxied success status, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "\"version_label\":\"26.2.20.08\"") {
		t.Fatalf("expected proxied create release response, got %s", rec.Body.String())
	}
}

func TestCreateDeploymentProxyUsesUpgradeService(t *testing.T) {
	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/deployments" {
			t.Fatalf("expected /v1/deployments, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		_, _ = w.Write([]byte(`{"id":"deployment-1","environment":"compose-stage","status":"deployed"}`))
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/deployments/create",
		bytes.NewBufferString(`{"target_version_id":"release-1","environment":"compose-stage","status":"deployed"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK && rec.Code != http.StatusCreated {
		t.Fatalf("expected proxied success status, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "\"environment\":\"compose-stage\"") {
		t.Fatalf("expected proxied create deployment response, got %s", rec.Body.String())
	}
}

func TestPromoteReleaseOrchestratesDeploymentAndSwitch(t *testing.T) {
	var deploymentCalled atomic.Bool
	var switchCalled atomic.Bool

	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/v1/deployments" && r.Method == http.MethodPost:
			deploymentCalled.Store(true)
			_, _ = w.Write([]byte(`{"id":"deployment-2","environment":"compose-promote","status":"deployed"}`))
		case r.URL.Path == "/v1/switches" && r.Method == http.MethodPost:
			if !deploymentCalled.Load() {
				t.Fatalf("switch was called before deployment")
			}
			switchCalled.Store(true)
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"target_type":"client","latest_version":"26.2.20.09"}`))
		default:
			t.Fatalf("unexpected call %s %s", r.Method, r.URL.Path)
		}
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/releases/promote",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","target_version_id":"release-2","environment":"compose-promote","operator":"service-manager"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected promote status 201, got %d body=%s", rec.Code, rec.Body.String())
	}
	if !deploymentCalled.Load() || !switchCalled.Load() {
		t.Fatalf("expected promote to call deployment and switch, got body=%s", rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "\"latest_version\":\"26.2.20.09\"") {
		t.Fatalf("expected promote response to include switched target, got %s", rec.Body.String())
	}
}

func TestMutatingReleaseActionsConflictWhenLocked(t *testing.T) {
	entered := make(chan struct{}, 1)
	release := make(chan struct{})

	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/switches" || r.Method != http.MethodPost {
			t.Fatalf("unexpected call %s %s", r.Method, r.URL.Path)
		}
		entered <- struct{}{}
		<-release
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"target_type":"client","latest_version":"26.2.20.10"}`))
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	firstDone := make(chan *httptest.ResponseRecorder, 1)
	go func() {
		req := httptest.NewRequest(
			http.MethodPost,
			"/v1/releases/switch",
			bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","to_version_id":"rv-1","operator":"service-manager"}`),
		)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		firstDone <- rec
	}()

	<-entered
	secondReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/releases/switch",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","to_version_id":"rv-2","operator":"service-manager"}`),
	)
	secondReq.Header.Set("Content-Type", "application/json")
	secondRec := httptest.NewRecorder()
	router.ServeHTTP(secondRec, secondReq)

	if secondRec.Code != http.StatusConflict {
		t.Fatalf("expected lock conflict status 409, got %d body=%s", secondRec.Code, secondRec.Body.String())
	}

	close(release)
	select {
	case firstRec := <-firstDone:
		if firstRec.Code != http.StatusCreated {
			t.Fatalf("expected first request to succeed, got %d body=%s", firstRec.Code, firstRec.Body.String())
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for first request")
	}
}

func TestCurrentReleaseOperationVisibleWhileRunning(t *testing.T) {
	entered := make(chan struct{}, 1)
	release := make(chan struct{})

	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/switches" || r.Method != http.MethodPost {
			t.Fatalf("unexpected call %s %s", r.Method, r.URL.Path)
		}
		entered <- struct{}{}
		<-release
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"target_type":"client","latest_version":"26.2.20.10"}`))
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	done := make(chan struct{}, 1)
	go func() {
		req := httptest.NewRequest(
			http.MethodPost,
			"/v1/releases/switch",
			bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","to_version_id":"rv-1","operator":"service-manager"}`),
		)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)
		done <- struct{}{}
	}()

	<-entered
	statusReq := httptest.NewRequest(http.MethodGet, "/v1/releases/operations/current", nil)
	statusRec := httptest.NewRecorder()
	router.ServeHTTP(statusRec, statusReq)

	if statusRec.Code != http.StatusOK {
		t.Fatalf("expected current operation status 200, got %d body=%s", statusRec.Code, statusRec.Body.String())
	}
	if !strings.Contains(statusRec.Body.String(), "\"operation\":\"switch\"") || !strings.Contains(statusRec.Body.String(), "\"status\":\"running\"") {
		t.Fatalf("expected running switch operation in current status, got %s", statusRec.Body.String())
	}

	close(release)
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for switch request to finish")
	}
}

func TestReleaseOperationHistoryRecordsCompletedPromote(t *testing.T) {
	upgradeService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/v1/deployments" && r.Method == http.MethodPost:
			_, _ = w.Write([]byte(`{"id":"deployment-3","environment":"compose-promote","status":"deployed"}`))
		case r.URL.Path == "/v1/switches" && r.Method == http.MethodPost:
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"target_type":"client","latest_version":"26.2.20.11"}`))
		default:
			t.Fatalf("unexpected call %s %s", r.Method, r.URL.Path)
		}
	}))
	defer upgradeService.Close()

	manager := runtime.NewManager([]runtime.ConfigService{
		{Name: "upgrade-service", Command: "sleep 30", WorkDir: ".", Address: upgradeService.URL},
	}, "local")
	router := NewRouterWithManager(manager)

	promoteReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/releases/promote",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","target_version_id":"release-3","environment":"compose-promote","operator":"service-manager"}`),
	)
	promoteReq.Header.Set("Content-Type", "application/json")
	promoteRec := httptest.NewRecorder()
	router.ServeHTTP(promoteRec, promoteReq)

	if promoteRec.Code != http.StatusCreated {
		t.Fatalf("expected promote status 201, got %d body=%s", promoteRec.Code, promoteRec.Body.String())
	}

	historyReq := httptest.NewRequest(http.MethodGet, "/v1/releases/operations/history", nil)
	historyRec := httptest.NewRecorder()
	router.ServeHTTP(historyRec, historyReq)

	if historyRec.Code != http.StatusOK {
		t.Fatalf("expected operation history status 200, got %d body=%s", historyRec.Code, historyRec.Body.String())
	}
	if !strings.Contains(historyRec.Body.String(), "\"operation\":\"promote\"") || !strings.Contains(historyRec.Body.String(), "\"status\":\"completed\"") {
		t.Fatalf("expected completed promote operation in history, got %s", historyRec.Body.String())
	}
	if strings.Contains(historyRec.Body.String(), "\"status\":\"running\"") {
		t.Fatalf("expected completed promote history without stale running record, got %s", historyRec.Body.String())
	}
}
