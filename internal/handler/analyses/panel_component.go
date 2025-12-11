package analyses

import (
	"log/slog"
	"net/http"
	"pet_checkup/internal/common"
	"pet_checkup/internal/repository/analyses/panel_component"
)

type PanelComponentHandler struct {
	Queries *panel_component.Queries
}

func NewPanelComponentHandler(queries *panel_component.Queries) *PanelComponentHandler {
	return &PanelComponentHandler{
		Queries: queries,
	}
}

func (h *PanelComponentHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newPanelComponent, err := common.Decode[panel_component.CreateParams](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		panelComponent, repoErr := h.Queries.Create(r.Context(), newPanelComponent)
		if repoErr != nil {
			http.Error(w, repoErr.Error(), http.StatusInternalServerError)
			return
		}
		err = common.Encode(w, r, http.StatusCreated, panelComponent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error("Encoding Err", "err", err.Error())
			return
		}
	})
}

func (h *PanelComponentHandler) UpdateByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := common.ExtractID(r)
		updateComponent, err := common.Decode[panel_component.UpdateByIDParams](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updateComponent.ID = id
		updated, repoErr := h.Queries.UpdateByID(r.Context(), updateComponent)
		if repoErr != nil {
			http.Error(w, repoErr.Error(), http.StatusInternalServerError)
			return
		}
		err = common.Encode(w, r, http.StatusOK, updated)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error("Encoding Err", "err", err.Error())
			return
		}
	})
}

func (h *PanelComponentHandler) DeleteByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

func (h *PanelComponentHandler) GetByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := common.ExtractID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		panelComponent, repoErr := h.Queries.GetByID(r.Context(), id)
		if repoErr != nil {
			http.Error(w, repoErr.Error(), http.StatusInternalServerError)
			return
		}
		err = common.Encode(w, r, http.StatusOK, panelComponent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error("Encoding Err", "err", err.Error())
			return
		}
	})
}
