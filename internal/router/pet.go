package router

import (
	"net/http"
	handler "pet_checkup/internal/handler/pet"
	db "pet_checkup/internal/repository/pet"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PetRouter(DbPool *pgxpool.Pool) http.Handler {
	queries := db.New(DbPool)
	petHandler := handler.NewPetHandler(queries)

	router := MakeCRUDRouter(RouteConfig{
		BasePath: "/pets",
		Handler:  petHandler,
	})
	router.Handle("POST /pets/{id}/weight", petHandler.AddWeight())
	router.Handle("GET /pets/{id}/weight", petHandler.GetPetWeights())
	router.Handle("DELETE /pets/{id}/weight", petHandler.DeleteWeight())
	return router
}
