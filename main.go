package main

import (
	"log"
	"net/http"

	"pet_checkup/languages"
	"pet_checkup/panels"
	"pet_checkup/pets"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}
	baseRouter := chi.NewRouter()
	baseRouter.Use(middleware.Recoverer)
	apiRouter := chi.NewRouter()
	baseRouter.Mount("/api", apiRouter)

	apiRouter.Mount("/languages", languages.LanguageRouter())
	apiRouter.Mount("/pets", pets.PetRouter())
	apiRouter.Mount("/panels", panels.PanelsRouter())
	http.ListenAndServe(":8000", baseRouter)
}
