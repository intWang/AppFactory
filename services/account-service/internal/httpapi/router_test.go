package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"appfactory/account-service/internal/storage"
)

func TestProvidersEndpointReturnsConfiguredProviders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/providers", nil)
	rec := httptest.NewRecorder()

	NewRouterWithDependencies(storage.NewMemoryStore(), defaultProviders()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "\"google\"") || !strings.Contains(body, "\"wechat\"") {
		t.Fatalf("expected provider list in response body, got %s", body)
	}
}

func TestRegisterEndpointReturnsCreatedAccountFromRequestBody(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/accounts/register",
		bytes.NewBufferString(`{"email":"new@example.com","password":"secret123","nickname":"New User"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	NewRouterWithDependencies(storage.NewMemoryStore(), defaultProviders()).ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "\"new@example.com\"") || !strings.Contains(body, "\"registered\"") {
		t.Fatalf("expected created account payload in response body, got %s", body)
	}
}
