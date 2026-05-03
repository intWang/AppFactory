package httpapi

import (
	"encoding/json"
	"net/http"

	"appfactory/account-service/internal/domain"
	"appfactory/account-service/internal/storage"
	sharedhealth "appfactory/shared-go/health"
	"appfactory/shared-go/httpx"
)

func NewRouter() http.Handler {
	store := storage.NewMemoryStore()
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, sharedhealth.Snapshot{
			Service: "account-service",
			Status:  "ok",
			Checks: map[string]string{
				"http": "ok",
			},
		})
	})

	mux.HandleFunc("/v1/accounts/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req domain.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		user := domain.User{
			ID:       "user-local-created",
			Email:    req.Email,
			Nickname: req.Nickname,
		}
		httpx.WriteJSON(w, http.StatusCreated, domain.RegisterResponse{
			Status: "registered",
			User:   user,
		})
	})
	mux.HandleFunc("/v1/accounts/login", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "local account login placeholder",
		})
	})
	mux.HandleFunc("/v1/accounts/logout", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"message": "logout placeholder",
		})
	})
	mux.HandleFunc("/v1/accounts/me", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, store.CurrentUser)
	})
	mux.HandleFunc("/v1/providers", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"providers": store.Providers,
		})
	})

	return mux
}
