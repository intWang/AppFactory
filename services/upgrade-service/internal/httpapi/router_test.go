package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"appfactory/upgrade-service/internal/storage"
)

func TestCheckClientUsesRequestBuildForUpgradeDecision(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/upgrade/check-client",
		bytes.NewBufferString(`{"product_slug":"shared-client","current_version":"26.2.20.01","current_build":1}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	NewRouterWithRepository(storage.NewMemoryStore()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "\"upgrade_mode\":\"optional\"") {
		t.Fatalf("expected optional upgrade response, got %s", body)
	}
}
