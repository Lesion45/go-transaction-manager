package main

import (
	"account-management-service/config"
	"account-management-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type App struct {
	server *http.Server
	router *gin.Engine
	// storage *postgres.Storage
	log *slog.Logger
}

func main() {
	cfg := config.MustLoad()

	log := logger.New(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("address", cfg.Server.Addr))
	log.Debug("logger debug mode enabled")

	// * Storage part starts
	// * Storage part ends

	router := gin.Default()

	server := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      router,
		ReadTimeout:  cfg.Server.TimeOut,
		WriteTimeout: cfg.Server.TimeOut,
		IdleTimeout:  cfg.Server.IdleTimeOut,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")
}
