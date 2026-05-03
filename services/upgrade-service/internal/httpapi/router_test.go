package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"appfactory/upgrade-service/internal/domain"
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

func TestReleaseLifecycleEndpoints(t *testing.T) {
	store := storage.NewMemoryStore()
	router := NewRouterWithRepository(store)

	releaseRec := httptest.NewRecorder()
	releaseReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/releases",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","version_label":"26.2.20.05","build_number":5,"upgrade_url":"https://example.com/client/26.2.20.05"}`),
	)
	releaseReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(releaseRec, releaseReq)
	if releaseRec.Code != http.StatusCreated {
		t.Fatalf("expected release status 201, got %d body=%s", releaseRec.Code, releaseRec.Body.String())
	}

	var release domain.ReleaseVersion
	if err := json.Unmarshal(releaseRec.Body.Bytes(), &release); err != nil {
		t.Fatalf("unmarshal release: %v", err)
	}
	if release.ID == "" {
		t.Fatalf("expected created release id, got empty")
	}

	deploymentRec := httptest.NewRecorder()
	deploymentReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/deployments",
		bytes.NewBufferString(`{"target_version_id":"`+release.ID+`","environment":"local","status":"deployed"}`),
	)
	deploymentReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(deploymentRec, deploymentReq)
	if deploymentRec.Code != http.StatusCreated {
		t.Fatalf("expected deployment status 201, got %d body=%s", deploymentRec.Code, deploymentRec.Body.String())
	}

	switchRec := httptest.NewRecorder()
	switchReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/switches",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","to_version_id":"`+release.ID+`","operator":"am"}`),
	)
	switchReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(switchRec, switchReq)
	if switchRec.Code != http.StatusCreated {
		t.Fatalf("expected switch status 201, got %d body=%s", switchRec.Code, switchRec.Body.String())
	}
	if !strings.Contains(switchRec.Body.String(), "\"latest_version\":\"26.2.20.05\"") {
		t.Fatalf("expected switched target to latest version 26.2.20.05, got %s", switchRec.Body.String())
	}

	activeRec := httptest.NewRecorder()
	activeReq := httptest.NewRequest(http.MethodGet, "/v1/targets/active", nil)
	router.ServeHTTP(activeRec, activeReq)
	if activeRec.Code != http.StatusOK {
		t.Fatalf("expected active targets status 200, got %d body=%s", activeRec.Code, activeRec.Body.String())
	}
	if !strings.Contains(activeRec.Body.String(), "\"latest_version\":\"26.2.20.05\"") {
		t.Fatalf("expected active targets to include switched version, got %s", activeRec.Body.String())
	}

	rollbackRec := httptest.NewRecorder()
	rollbackReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/rollbacks",
		bytes.NewBufferString(`{"product_slug":"shared-client","target_type":"client","rolled_back_to_version_id":"release-client-4","operator":"qa"}`),
	)
	rollbackReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rollbackRec, rollbackReq)
	if rollbackRec.Code != http.StatusCreated {
		t.Fatalf("expected rollback status 201, got %d body=%s", rollbackRec.Code, rollbackRec.Body.String())
	}
	if !strings.Contains(rollbackRec.Body.String(), "\"latest_version\":\"26.2.20.04\"") {
		t.Fatalf("expected rollback target to latest version 26.2.20.04, got %s", rollbackRec.Body.String())
	}
}
