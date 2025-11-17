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

	return MakeCRUDRouter(RouteConfig{
		BasePath: "/pets",
		Handler:  petHandler,
	})
}
