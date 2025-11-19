package router

import (
	"net/http"
	"pet_checkup/internal/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
)

// CRUDHandler defines the common CRUD operations that handlers can implement.
// Handlers can implement all or a subset of these methods.
type CRUDHandler any

type ListHandler interface {
	CRUDHandler
	GetMany() http.Handler
}

type CreateHandler interface {
	CRUDHandler
	Create() http.Handler
}

type GetByIDHandler interface {
	CRUDHandler
	GetByID() http.Handler
}

type UpdateByIDHandler interface {
	CRUDHandler
	UpdateByID() http.Handler
}

type DeleteByIDHandler interface {
	CRUDHandler
	DeleteByID() http.Handler
}

type RouteConfig struct {
	BasePath string
	Handler  CRUDHandler
}

// MakeCRUDRouter creates a generic CRUD router based on which interfaces the handler implements
func MakeCRUDRouter(config RouteConfig) *http.ServeMux {
	mux := http.NewServeMux()
	basePath := config.BasePath
	handler := config.Handler

	if h, ok := handler.(ListHandler); ok {
		mux.Handle("GET "+basePath, h.GetMany())
	}

	if h, ok := handler.(CreateHandler); ok {
		mux.Handle("POST "+basePath, h.Create())
	}

	if h, ok := handler.(GetByIDHandler); ok {
		mux.Handle("GET "+basePath+"/{id}", h.GetByID())
	}

	if h, ok := handler.(UpdateByIDHandler); ok {
		mux.Handle("PUT "+basePath+"/{id}", h.UpdateByID())
	}

	if h, ok := handler.(DeleteByIDHandler); ok {
		mux.Handle("DELETE "+basePath+"/{id}", h.DeleteByID())
	}

	return mux
}

func BaseRouter(DbPool *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	apiMux := http.NewServeMux()
	apiMux.Handle("/", PetRouter(DbPool))
	apiMux.Handle("/analyses/", PanelRouter(DbPool))
	apiMux.Handle("/analyses/", ComponentRouter(DbPool))

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	return middleware.Chain(
		mux,
		middleware.StripTrailingSlash,
		middleware.LogRequest)
}
