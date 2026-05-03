package httpapi

import (
	"net/http"

	"appfactory/service-manager/internal/runtime"
	sharedhealth "appfactory/shared-go/health"
	"appfactory/shared-go/httpx"
)

func NewRouter() http.Handler {
	registry := runtime.NewRegistry()
	mux := http.NewServeMux()

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
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"services": registry.Services})
	})
	mux.HandleFunc("/v1/services/start", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"message": "start placeholder"})
	})
	mux.HandleFunc("/v1/services/stop", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"message": "stop placeholder"})
	})
	mux.HandleFunc("/v1/services/restart", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"message": "restart placeholder"})
	})
	mux.HandleFunc("/v1/services/health", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"services": registry.Services})
	})
	mux.HandleFunc("/v1/services/switch-profile", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"message": "switch-profile placeholder"})
	})
	mux.HandleFunc("/v1/services/targets", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"targets": registry.Services})
	})

	return mux
}
