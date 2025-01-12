package components

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ComponentsRouter(DBPool *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	h := ComponentsHandler{DBPool}
	return r.Route("/components", func(r chi.Router) {
		r.Post("/", h.CreateComponent)
		r.Get("/", h.GetComponents)

		r.Post("/measurement-unit", h.CreateMeasurementUnit)
		r.Put("/measurement-unit/{measurementUnitID}", h.UpdateMeasurementUnit)
		r.Get("/measurement-unit", h.GetMeasurementUnits)
	})
}
