package panels

import (
	"encoding/json"
	"net/http"
	"pet_checkup/api/utils"
	db "pet_checkup/internal/db/panels/measurement_units"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MeasurementUnitsHandler struct {
	DbPool *pgxpool.Pool
}

func (h *MeasurementUnitsHandler) CreateMeasurementUnit(w http.ResponseWriter, r *http.Request) {
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

func (h *MeasurementUnitsHandler) UpdateMeasurementUnit(w http.ResponseWriter, r *http.Request) {
	unitID := utils.IDfromUrl(r, MeasurementUnitID)
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

func (h *MeasurementUnitsHandler) GetMeasurementUnits(w http.ResponseWriter, r *http.Request) {
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
