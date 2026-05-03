package httpapi

import (
	"context"
	"encoding/json"
	"net/http"

	"appfactory/account-service/internal/domain"
	"appfactory/account-service/internal/storage"
	sharedhealth "appfactory/shared-go/health"
	"appfactory/shared-go/httpx"
)

func NewRouter() http.Handler {
	return NewRouterWithDependencies(storage.NewMemoryStore(), defaultProviders())
}

func NewRouterWithDependencies(store storage.Repository, providers []domain.Provider) http.Handler {
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
		user, err := store.Register(r.Context(), req)
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
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
		user, err := store.GetCurrentUser(r.Context())
		if err != nil {
			httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpx.WriteJSON(w, http.StatusOK, user)
	})
	mux.HandleFunc("/v1/providers", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{
			"providers": providers,
		})
	})

	return mux
}

func defaultProviders() []domain.Provider {
	return []domain.Provider{
		{Key: "google", Enabled: false, Available: true},
		{Key: "apple", Enabled: false, Available: true},
		{Key: "wechat", Enabled: false, Available: true},
	}
}

type Config struct {
	ServiceName      string   `yaml:"service_name"`
	HTTPPort         string   `yaml:"http_port"`
	Environment      string   `yaml:"environment"`
	PostgresDSN      string   `yaml:"postgres_dsn"`
	RedisAddr        string   `yaml:"redis_addr"`
	SessionMode      string   `yaml:"session_mode"`
	StorageMode      string   `yaml:"storage_mode"`
	ProviderRegistry []string `yaml:"provider_registry"`
}

func NewPostgresRouter(ctx context.Context, cfg Config) (http.Handler, func(), error) {
	store, err := storage.NewPostgresStore(ctx, cfg.PostgresDSN)
	if err != nil {
		return nil, nil, err
	}
	providers := make([]domain.Provider, 0, len(cfg.ProviderRegistry))
	for _, key := range cfg.ProviderRegistry {
		providers = append(providers, domain.Provider{Key: key, Available: true, Enabled: false})
	}
	return NewRouterWithDependencies(store, providers), func() {}, nil
}
