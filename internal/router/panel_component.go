package router

import (
	"net/http"
	"pet_checkup/internal/handler/analyses"
	"pet_checkup/internal/repository/analyses/panel_component"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PanelComponentRouter(DbPool *pgxpool.Pool) http.Handler {
	queries := panel_component.New(DbPool)
	handler := analyses.NewPanelComponentHandler(queries)

	router := MakeCRUDRouter(RouteConfig{
		BasePath: "/panel_component",
		Handler:  handler})
	return router
}
