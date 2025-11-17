package router

import "net/http"

// CRUDHandler defines the common CRUD operations that handlers can implement.
// Handlers can implement all or a subset of these methods.
type CRUDHandler any

// ListHandler provides a method to list all resources
type ListHandler interface {
	CRUDHandler
	GetAll() http.Handler
}

// CreateHandler provides a method to create a new resource
type CreateHandler interface {
	CRUDHandler
	Create() http.Handler
}

// GetByIDHandler provides a method to get a resource by ID
type GetByIDHandler interface {
	CRUDHandler
	GetByID() http.Handler
}

// UpdateByIDHandler provides a method to update a resource by ID
type UpdateByIDHandler interface {
	CRUDHandler
	UpdateByID() http.Handler
}

// DeleteByIDHandler provides a method to delete a resource by ID
type DeleteByIDHandler interface {
	CRUDHandler
	DeleteByID() http.Handler
}

// RouteConfig defines the configuration for a CRUD route
type RouteConfig struct {
	BasePath string
	Handler  CRUDHandler
}

// MakeCRUDRouter creates a generic CRUD router based on which interfaces the handler implements
func MakeCRUDRouter(config RouteConfig) http.Handler {
	mux := http.NewServeMux()
	basePath := config.BasePath
	handler := config.Handler

	// Register LIST endpoint if handler implements ListHandler
	if h, ok := handler.(ListHandler); ok {
		mux.Handle("GET "+basePath, h.GetAll())
	}

	// Register CREATE endpoint if handler implements CreateHandler
	if h, ok := handler.(CreateHandler); ok {
		mux.Handle("POST "+basePath, h.Create())
	}

	// Register GET by ID endpoint if handler implements GetByIDHandler
	if h, ok := handler.(GetByIDHandler); ok {
		mux.Handle("GET "+basePath+"/{id}", h.GetByID())
	}

	// Register UPDATE endpoint if handler implements UpdateByIDHandler
	if h, ok := handler.(UpdateByIDHandler); ok {
		mux.Handle("PUT "+basePath+"/{id}", h.UpdateByID())
	}

	// Register DELETE endpoint if handler implements DeleteByIDHandler
	if h, ok := handler.(DeleteByIDHandler); ok {
		mux.Handle("DELETE "+basePath+"/{id}", h.DeleteByID())
	}

	return mux
}
