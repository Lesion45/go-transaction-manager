package main

import (
	"github.com/gin-gonic/gin"
	"go-transaction-manager/config"
	"go-transaction-manager/pkg/logger"
	"log/slog"
	"net/http"
)

type App struct {
	srv    *http.Server
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
