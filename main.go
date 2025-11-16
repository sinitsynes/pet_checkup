package main

import (
	"context"
	"log/slog"
	"net/http"
	"pet_checkup/configuration"
	"pet_checkup/internal/router"
	"pet_checkup/platform/database"
	"pet_checkup/platform/logger"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.Error("failed to run", "error", err)
	}
}

func run(ctx context.Context) error {

	config := configuration.LoadConfig()
	logger.Setup(config.Server.Stand)

	DbPool, err := database.InitDB(ctx, config.DB.ConnectionString())
	if err != nil {
		slog.Error("Unable to establish Database connection", "error", err)
		return err
	}
	defer DbPool.Close()
	httpRouter := router.BaseRouter(DbPool)
	serverAddr := config.Server.Host + ":" + config.Server.Port
	slog.Info("Starting server", "address", serverAddr)
	serverErr := http.ListenAndServe(serverAddr, httpRouter)
	if serverErr != nil {
		slog.Error("failed to start server", "error", serverErr)
		return serverErr
	}
	return nil
}
