package httpapi

import (
	"encoding/json"
	"net/http"

	"appfactory/shared-go/httpx"
	sharedhealth "appfactory/shared-go/health"
	"appfactory/upgrade-service/internal/domain"
	"appfactory/upgrade-service/internal/storage"
)

func NewRouter() http.Handler {
	store := storage.NewMemoryStore()
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, sharedhealth.Snapshot{
			Service: "upgrade-service",
			Status:  "ok",
			Checks: map[string]string{
				"http": "ok",
			},
		})
	})

	mux.HandleFunc("/v1/upgrade/check-client", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req domain.CheckUpgradeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		mode := store.ClientTarget.UpgradeModeForBuild(req.CurrentBuild)
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"target":        store.ClientTarget,
			"upgrade_mode":  mode,
			"force_upgrade": mode == "forced",
		})
	})
	mux.HandleFunc("/v1/upgrade/check-service", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req domain.CheckUpgradeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		mode := store.ServiceTarget.UpgradeModeForBuild(req.CurrentBuild)
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"target":        store.ServiceTarget,
			"upgrade_mode":  mode,
			"force_upgrade": mode == "forced",
		})
	})
	mux.HandleFunc("/v1/releases", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusCreated, map[string]string{"message": "release placeholder"})
	})
	mux.HandleFunc("/v1/deployments", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusCreated, map[string]string{"message": "deployment placeholder"})
	})
	mux.HandleFunc("/v1/switches", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusCreated, map[string]string{"message": "switch placeholder"})
	})
	mux.HandleFunc("/v1/rollbacks", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusCreated, map[string]string{"message": "rollback placeholder"})
	})
	mux.HandleFunc("/v1/targets/active", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"client":  store.ClientTarget,
			"service": store.ServiceTarget,
		})
	})

	return mux
}
