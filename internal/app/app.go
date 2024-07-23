package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-transaction-manager/config"
	v1 "go-transaction-manager/internal/controller/http/v1"
	"go-transaction-manager/internal/repository"
	"go-transaction-manager/internal/service"
	"go-transaction-manager/pkg/logger"
	"go-transaction-manager/pkg/logger/sl"
	"go-transaction-manager/pkg/postgres"
	"log/slog"
	"net/http"
	"os"
)

func Run() {
	// Configuration
	cfg := config.MustLoad()

	// Logger
	log := logger.New(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("Initializing application...")
	log.Debug("logger debug mode enabled")

	// Repositories
	log.Info("Initializing postgres...")
	pg, err := postgres.NewPG(context.Background(), cfg.DB)
	if err != nil {
		log.Error("database initialization error", sl.Err(err))
		os.Exit(1)
	}

	err = pg.PostgresHealthCheck(context.Background())
	if err != nil {
		log.Error("Postgres doesn't response", sl.Err(err))
		os.Exit(1)
	}
	defer pg.DB.Close()
	log.Info("Initializing postgres: successful!")

	log.Info("Initializing repositories...")
	repositories := repository.NewRepositories(pg)
	log.Info("Initializing repositories: successful!")

	// Services dependencies
	log.Info("Initializing services...")
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps)
	log.Info("Initializing services: successful!")

	log.Info("Initializing server...")
	router := gin.Default()

	routes := v1.NewRouter(router, services, log)
	if routes == nil {
		log.Error("Initialization routes error", sl.Err(err))
		os.Exit(1)
	}

	server := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      router,
		ReadTimeout:  cfg.Server.TimeOut,
		WriteTimeout: cfg.Server.TimeOut,
		IdleTimeout:  cfg.Server.IdleTimeOut,
	}
	log.Info("Initializing server: OK!")

	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")
}
