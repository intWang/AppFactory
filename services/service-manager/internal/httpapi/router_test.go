package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
