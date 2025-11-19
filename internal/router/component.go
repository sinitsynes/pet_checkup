package router

import (
	"net/http"
	"pet_checkup/internal/handler/analyses"
	"pet_checkup/internal/repository/analyses/component"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ComponentRouter(DbPool *pgxpool.Pool) http.Handler {
	queries := component.New(DbPool)
	handler := analyses.NewComponentHandler(queries)

	router := MakeCRUDRouter(RouteConfig{
		BasePath: "/component",
		Handler:  handler})
	return router
}
