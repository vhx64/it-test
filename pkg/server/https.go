package server

import (
	"context"

	"it-test/config"

	"it-test/pkg/logs"

	"fmt"

	"io/ioutil"

	"net/http"

	"os"

	"os/signal"

	"strings"

	"syscall"

	"time"

	"github.com/ghodss/yaml"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/cors"

	"github.com/sirupsen/logrus"
)

const (
	teardownTimeout = 10
)

func RunHTTPServer(ctx context.Context, cfg *config.AppConfig, port string,
	createHandler func(router chi.Router) http.Handler) {
	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGINT, syscall.SIGTERM)

	startCtx, startCancel := context.WithCancel(ctx)
	apiRouter := chi.NewRouter()
	setMiddlewares(cfg, apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Get("/health", Health)
	rootRouter.Get("/api/v1/it-test/api-docs/swagger.json", Swagger)

	// we are mounting all APIs under /api path
	if createHandler != nil {
		rootRouter.Mount("/api/v1/it-test", createHandler(apiRouter))
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      rootRouter,
		WriteTimeout: cfg.ApplicationAPITimeout,
		ReadTimeout:  cfg.ApplicationAPITimeout,
	}

	go func() {
		<-osChan
		logrus.WithContext(startCtx).Debug("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), teardownTimeout*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logrus.WithContext(startCtx).
				WithError(err).
				Fatal("Could not gracefully shutdown the httpServer")
		}

		startCancel()
		logrus.WithContext(ctx).Info("Server stopped")
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithContext(startCtx).WithError(err).Fatal("Server shutting down, error occurred")
	}
}

func setMiddlewares(cfg *config.AppConfig, router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.ApplicationAPITimeout))
	addCorsMiddleware(router, cfg.CORSAllowedOrigins)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux, origins string) {
	allowedOrigins := strings.Split(origins, ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-App-Id"},
		ExposedHeaders:   []string{"Link", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}

// Health operation middleware
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Swagger operation middleware
func Swagger(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./api/api.yml")
	if err != nil {
		logrus.WithContext(r.Context()).WithError(err).Info("failed to read api.yml")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	j, err := yaml.YAMLToJSON(file)

	if err != nil {
		logrus.WithContext(r.Context()).WithError(err).Info("failed to convert yaml to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		logrus.WithContext(r.Context()).WithError(err).Info("failed to write api.yml to response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
