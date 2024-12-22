package main

import (
	"context"
	"net/http"
	"pet_checkup/api"
	"pet_checkup/internal/app"
	"pet_checkup/internal/db"
)

func main() {
	ctx := context.Background()
	app, err := run(ctx)
	if err != nil {
		panic(err.Error())
	}
	router := api.AppRouter(app)

	defer app.DbPool.Close()

	http.ListenAndServe(":8000", router)
}

func run(ctx context.Context) (app.Application, error) {
	app := app.Application{}
	connString, err := db.ConnectionString()
	if err != nil {
		return app, err
	}
	dbPool, err := db.NewDBPool(ctx, connString)
	if err != nil {
		return app, err
	}
	app.DbPool = dbPool
	return app, nil
}
