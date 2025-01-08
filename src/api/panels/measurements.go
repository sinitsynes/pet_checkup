package panels

import (
	"encoding/json"
	"net/http"
	"pet_checkup/api/utils"
	db "pet_checkup/internal/db/panels/measurements"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MeasurementsHandler struct {
	DbPool *pgxpool.Pool
}

func (h *MeasurementsHandler) CreateMeasurement(w http.ResponseWriter, r *http.Request) {
	petID := utils.IDfromUrl(r, utils.PetIDParam)
	newMeasurementParams := []db.CreateMeasurementParams{}
	decoderErr := json.NewDecoder(r.Body).Decode(&newMeasurementParams)
	if decoderErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i, item := range newMeasurementParams {
		item.PetID.Int32 = petID
		item.PetID.Valid = true
		newMeasurementParams[i] = item
	}
	query := db.New(h.DbPool)
	createdMeasurement, dbErr := query.CreateMeasurement(r.Context(), newMeasurementParams)
	if dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(dbErr.Error()))
		return
	}
	response, jsonErr := json.Marshal(createdMeasurement)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
