package httpapi

import (
	"context"
	"encoding/json"
	"net/http"

	sharedhealth "appfactory/shared-go/health"
	"appfactory/shared-go/httpx"
	"appfactory/upgrade-service/internal/domain"
	"appfactory/upgrade-service/internal/storage"
)

func NewRouter() http.Handler {
	return NewRouterWithRepository(storage.NewMemoryStore())
}

func NewRouterWithRepository(store storage.Repository) http.Handler {
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
		target, err := store.GetTarget(r.Context(), "client")
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		target.CurrentVersion = req.CurrentVersion
		target.CurrentBuild = req.CurrentBuild
		mode := target.UpgradeModeForBuild(req.CurrentBuild)
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"target":        target,
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
		target, err := store.GetTarget(r.Context(), "service")
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		target.CurrentVersion = req.CurrentVersion
		target.CurrentBuild = req.CurrentBuild
		mode := target.UpgradeModeForBuild(req.CurrentBuild)
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"target":        target,
			"upgrade_mode":  mode,
			"force_upgrade": mode == "forced",
		})
	})
	mux.HandleFunc("/v1/releases", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req domain.CreateReleaseRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		release, err := store.CreateRelease(r.Context(), req)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, release)
	})
	mux.HandleFunc("/v1/deployments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req domain.CreateDeploymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		deployment, err := store.CreateDeployment(r.Context(), req)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, deployment)
	})
	mux.HandleFunc("/v1/switches", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req domain.SwitchTargetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		target, err := store.SwitchTarget(r.Context(), req)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, target)
	})
	mux.HandleFunc("/v1/rollbacks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req domain.RollbackTargetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		target, err := store.RollbackTarget(r.Context(), req)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusCreated, target)
	})
	mux.HandleFunc("/v1/targets/active", func(w http.ResponseWriter, r *http.Request) {
		targets, err := store.GetActiveTargets(r.Context())
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusOK, targets)
	})

	return mux
}

type Config struct {
	ServiceName           string `yaml:"service_name"`
	HTTPPort              string `yaml:"http_port"`
	Environment           string `yaml:"environment"`
	PostgresDSN           string `yaml:"postgres_dsn"`
	RedisAddr             string `yaml:"redis_addr"`
	ForcedUpgradeBuildGap int    `yaml:"forced_upgrade_build_gap"`
	VersionFormat         string `yaml:"version_format"`
	StorageMode           string `yaml:"storage_mode"`
}

func NewPostgresRouter(ctx context.Context, cfg Config) (http.Handler, func(), error) {
	store, err := storage.NewPostgresStore(ctx, cfg.PostgresDSN, cfg.ForcedUpgradeBuildGap)
	if err != nil {
		return nil, nil, err
	}
	return NewRouterWithRepository(store), func() {}, nil
}
