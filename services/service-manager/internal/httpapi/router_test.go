package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

	router := NewRouterWithRegistry([]ServiceRuntime{
		{Name: "account-service", Address: healthyService.URL, Status: "registered", Profile: "local"},
	})

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
