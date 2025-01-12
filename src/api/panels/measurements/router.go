package measurements

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MeasurementsRouter(DBPool *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	h := MeasurementsHandler{DBPool}
	return r.Route("/weight", func(r chi.Router) {
		r.Post("/measurements/{petID}", h.CreateMeasurement)
	})
}
