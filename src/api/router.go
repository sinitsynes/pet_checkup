package api

import (
	"pet_checkup/api/languages"
	"pet_checkup/api/panels/components"
	"pet_checkup/api/panels/measurements"
	"pet_checkup/api/panels/weight"
	"pet_checkup/api/pets"
	"pet_checkup/internal/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AppRouter(app *app.Application) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	router.Route("/api", func(r chi.Router) {
		languages.LanguageRoute(app.DbPool)
		pets.PetsRoute(app.DbPool)
		weight.WeightRouter(app.DbPool)
		components.ComponentsRouter(app.DbPool)
		measurements.MeasurementsRouter(app.DbPool)
	})
	return router
}
