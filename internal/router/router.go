package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func BaseRouter(DbPool *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	apiMux := http.NewServeMux()
	apiMux.Handle("/", PetRouter(DbPool))

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))
	return mux
}
