package router

import (
	"net/http"
	"pet_checkup/internal/handler/analyses"
	"pet_checkup/internal/repository/analyses/panel"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PanelRouter(DbPool *pgxpool.Pool) http.Handler {
	queries := panel.New(DbPool)
	panelHandler := analyses.NewPanelHandler(queries)

	router := MakeCRUDRouter(RouteConfig{
		BasePath: "/panel",
		Handler:  panelHandler,
	})

	return router
}
