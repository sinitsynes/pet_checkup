package panel

import (
	"log/slog"
	"net/http"
	"pet_checkup/internal/common"
	"pet_checkup/internal/repository/analyses/panel"
)

type PanelHandler struct {
	Queries *panel.Queries
}

func NewPanelHandler(queries *panel.Queries) *PanelHandler {
	return &PanelHandler{
		Queries: queries,
	}
}

func (h *PanelHandler) GetByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := common.ExtractID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		panel, repoErr := h.Queries.GetByID(r.Context(), id)
		if repoErr != nil {
			http.Error(w, repoErr.Error(), http.StatusInternalServerError)
			return
		}
		err = common.Encode(w, r, http.StatusOK, panel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error("Encoding Err", "err", err.Error())
			return
		}
	})
}

func (h *PanelHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newPanel, err := common.Decode[panel.CreateParams](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		panel, repoErr := h.Queries.Create(r.Context(), newPanel)
		if repoErr != nil {
			http.Error(w, repoErr.Error(), http.StatusInternalServerError)
			return
		}
		err = common.Encode(w, r, http.StatusCreated, panel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error("Encoding Err", "err", err.Error())
			return
		}
	})
}

func (h *PanelHandler) UpdateByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := common.ExtractID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatePanel, err := common.Decode[panel.UpdateByIDParams](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatePanel.ID = id
		updated, repoErr := h.Queries.UpdateByID(r.Context(), updatePanel)
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

func (h *PanelHandler) DeleteByID() http.Handler {
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
