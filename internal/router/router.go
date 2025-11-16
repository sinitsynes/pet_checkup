package router

import (
	"net/http"
	"pet_checkup/internal/router/language"

	"github.com/jackc/pgx/v5/pgxpool"
)

func BaseRouter(DbPool *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", language.LanguageRouter(DbPool)))
	return mux
}
