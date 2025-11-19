package analyses

import (
	"log/slog"
	"net/http"
	"pet_checkup/internal/common"
	"pet_checkup/internal/repository/analyses/component"
)

type ComponentHandler struct {
	Queries *component.Queries
}

func NewComponentHandler(queries *component.Queries) *ComponentHandler {
	return &ComponentHandler{
		Queries: queries,
	}
}

func (h *ComponentHandler) Create() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			newComponent, err := common.Decode[component.CreateParams](r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			component, repoErr := h.Queries.Create(r.Context(), newComponent)
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			err = common.Encode(w, r, http.StatusCreated, component)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (h *ComponentHandler) GetByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := common.ExtractID(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			component, repoErr := h.Queries.GetByID(r.Context(), id)
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			err = common.Encode(w, r, http.StatusOK, component)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (h *ComponentHandler) UpdateByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := common.ExtractID(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updateComponent, err := common.Decode[component.UpdateByIDParams](r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updateComponent.ID = id
			component, repoErr := h.Queries.UpdateByID(r.Context(), updateComponent)
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			err = common.Encode(w, r, http.StatusOK, component)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (h *ComponentHandler) DeleteByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := common.ExtractID(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			repoErr := h.Queries.DeleteByID(r.Context(), id)
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
}
