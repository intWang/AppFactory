package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProvidersEndpointReturnsConfiguredProviders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/providers", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "\"google\"") || !strings.Contains(body, "\"wechat\"") {
		t.Fatalf("expected provider list in response body, got %s", body)
	}
}
