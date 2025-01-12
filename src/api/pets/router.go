package pets

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func PetsRoute(DBPool *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	h := PetHandler{DBPool}
	return r.Route("/pets", func(r chi.Router) {
		r.Post("/", h.CreateRecord)
		r.Put("/{petID}", h.EditRecord)
		r.Get("/{petID}", h.GetRecord)
		r.Delete("/{petID}", h.DeleteRecord)
	})
}
