package api

import (
	"pet_checkup/api/language"
	"pet_checkup/api/panels"
	"pet_checkup/api/pets"
	"pet_checkup/internal/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AppRouter(app app.Application) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	languageHandler := language.LanguageHandler{DbPool: app.DbPool}
	petsHandler := pets.PetHandler{DbPool: app.DbPool}

	panelComponentHandler := panels.ComponentHandler{DbPool: app.DbPool}
	panelMeasurementUnitHandler := panels.MeasurementUnitsHandler{DbPool: app.DbPool}
	panelMeasurementHandler := panels.MeasurementsHandler{DbPool: app.DbPool}
	weightHandler := panels.WeightHandler{DbPool: app.DbPool}

	router.Route("/api", func(r chi.Router) {
		r.Route("/languages", func(r chi.Router) {
			r.Get("/", languageHandler.GetAvailableLanguages)
			r.Post("/", languageHandler.CreateLanguage)
		})
		r.Route("/pets", func(r chi.Router) {
			r.Post("/", petsHandler.CreateRecord)
			r.Put("/{petID}", petsHandler.EditRecord)
			r.Get("/{petID}", petsHandler.GetRecord)
			r.Delete("/{petID}", petsHandler.DeleteRecord)
		})
		r.Route("/panels", func(r chi.Router) {
			r.Post("/weight", weightHandler.AddWeightRecord)
			r.Get("/weight/{petID}", weightHandler.GetWeightRecords)
			r.Delete("/weight/{weightID}", weightHandler.DeleteWeightRecord)

			r.Post("/components/measurement-unit", panelMeasurementUnitHandler.CreateMeasurementUnit)
			r.Put("/components/measurement-unit/{measurementUnitID}", panelMeasurementUnitHandler.UpdateMeasurementUnit)
			r.Get("/components/measurement-unit", panelMeasurementUnitHandler.GetMeasurementUnits)

			r.Post("/components", panelComponentHandler.CreateComponent)
			r.Get("/components", panelComponentHandler.GetComponents)

			r.Post("/measurements/{petID}", panelMeasurementHandler.CreateMeasurement)
		})
	})
	return router
}
