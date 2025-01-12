package languages

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LanguageRoute(DBPool *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	h := LanguagesHandler{DBPool}
	return r.Route("/languages", func(r chi.Router) {
		r.Get("/", h.GetAvailableLanguages)
		r.Post("/", h.CreateLanguage)
	})
}
