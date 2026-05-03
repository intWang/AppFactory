package httpapi

import (
	"net/http"

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
		httpx.WriteJSON(w, http.StatusCreated, map[string]any{
			"message": "local account registration placeholder",
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
