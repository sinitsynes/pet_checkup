package main

import (
	"context"
	"log/slog"
	"net/http"
	"pet_checkup/api"
	"pet_checkup/internal/app"
	"pet_checkup/internal/db"
)

func main() {
	app := run()
	router := api.AppRouter(app)

	defer app.DbPool.Close()

	http.ListenAndServe(":8000", router)
}

func run() *app.Application {
	logger := app.LoggerSetup()
	slog.SetDefault(logger)

	ctx := context.Background()

	application := app.ApplicationSetup()

	connString := application.Config.DB.ConnectionString

	dbPool, err := db.NewDBPool(ctx, connString)
	if err != nil {
		panic(err.Error())
	}
	application.DbPool = dbPool
	return application
}
