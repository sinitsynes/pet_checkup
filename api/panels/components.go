package panels

import (
	"encoding/json"
	"net/http"
	db "pet_checkup/internal/db/panels/components"

	"github.com/jackc/pgx/v5/pgxpool"
)

const MeasurementUnitID = "measurementUnitID"

type ComponentHandler struct {
	DbPool *pgxpool.Pool
}

func (h *ComponentHandler) CreateComponent(w http.ResponseWriter, r *http.Request) {
	newComponent := &db.CreatePanelComponentParams{}
	decoderErr := json.NewDecoder(r.Body).Decode(newComponent)
	if decoderErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	query := db.New(h.DbPool)
	_, dbErr := query.CreatePanelComponent(r.Context(), *newComponent)
	if dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ComponentHandler) GetComponents(w http.ResponseWriter, r *http.Request) {
	langLabel := r.URL.Query().Get("lang")
	q := db.New(h.DbPool)
	components, dbErr := q.GetComponents(r.Context(), langLabel)
	if dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, jsonErr := json.Marshal(components)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
