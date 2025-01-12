package components

import (
	"encoding/json"
	"net/http"
	"pet_checkup/api/utils"
	db "pet_checkup/internal/db/panels/measurement_units"

	"github.com/go-chi/chi/v5"
)

func (h *ComponentsHandler) CreateMeasurementUnit(w http.ResponseWriter, r *http.Request) {
	newMeasurementUnit := &db.CreateMeasurementUnitParams{}
	decoderErr := json.NewDecoder(r.Body).Decode(newMeasurementUnit)
	if decoderErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	query := db.New(h.DbPool)
	if dbErr := query.CreateMeasurementUnit(r.Context(), *newMeasurementUnit); dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ComponentsHandler) UpdateMeasurementUnit(w http.ResponseWriter, r *http.Request) {
	unitID, err := utils.IDfromUrl(r, "MeasurementUnitID")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	updateUnitParams := &db.UpdateMeasurementUnitParams{ID: unitID}
	decoderErr := json.NewDecoder(r.Body).Decode(updateUnitParams)
	if decoderErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := db.New(h.DbPool)
	if db_err := query.UpdateMeasurementUnit(r.Context(), *updateUnitParams); db_err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (h *ComponentsHandler) GetMeasurementUnits(w http.ResponseWriter, r *http.Request) {
	langLabel := chi.URLParam(r, "lang")
	query := db.New(h.DbPool)
	units, dbErr := query.GetMeasurementUnits(r.Context(), langLabel)
	if dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	response, jsonErr := json.Marshal(units)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
