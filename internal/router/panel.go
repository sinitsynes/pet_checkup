package router

import (
	"net/http"
	handler "pet_checkup/internal/handler/panel"
	db "pet_checkup/internal/repository/analyses/panel"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPanelRouter(DbPool *pgxpool.Pool) http.Handler {
	queries := db.New(DbPool)
	panelHandler := handler.NewPanelHandler(queries)

	router := MakeCRUDRouter(RouteConfig{
		BasePath: "/panel",
		Handler:  panelHandler,
	})

	return router
}
