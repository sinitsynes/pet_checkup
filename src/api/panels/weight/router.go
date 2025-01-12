package weight

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func WeightRouter(DBPool *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	h := WeightHandler{DBPool}
	return r.Route("/weight", func(r chi.Router) {
		r.Post("/", h.AddWeightRecord)
		r.Get("/{petID}", h.GetWeightRecords)
		r.Delete("/{weightID}", h.DeleteWeightRecord)
	})
}
