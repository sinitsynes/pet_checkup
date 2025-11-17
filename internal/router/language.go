package router

import (
	"net/http"
	handler "pet_checkup/internal/handler/language"
	db "pet_checkup/internal/repository/language"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LanguageRouter(DBPool *pgxpool.Pool) http.Handler {
	queries := db.New(DBPool)
	languageHandler := handler.NewLanguageHandler(queries)

	return MakeCRUDRouter(RouteConfig{
		BasePath: "/languages",
		Handler:  languageHandler,
	})
}
