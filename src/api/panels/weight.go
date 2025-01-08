package panels

import (
	"encoding/json"
	"net/http"
	"pet_checkup/api/utils"
	db "pet_checkup/internal/db/panels/weight"

	"github.com/jackc/pgx/v5/pgxpool"
)

const weightID = "weightID"

type WeightHandler struct {
	DbPool *pgxpool.Pool
}

func (handler WeightHandler) AddWeightRecord(w http.ResponseWriter, r *http.Request) {
	addWeightParams := &db.CreateWeightMeasureParams{}
	if decoderErr := json.NewDecoder(r.Body).Decode(addWeightParams); decoderErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	query := db.New(handler.DbPool)
	if dbErr := query.CreateWeightMeasure(r.Context(), *addWeightParams); dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler WeightHandler) GetWeightRecords(w http.ResponseWriter, r *http.Request) {
	petID := utils.IDfromUrl(r, "petID")
	query := db.New(handler.DbPool)
	measures, dbErr := query.GetWeightMeasures(r.Context(), petID)
	if dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(measures)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (handler WeightHandler) DeleteWeightRecord(w http.ResponseWriter, r *http.Request) {
	weightID := utils.IDfromUrl(r, weightID)
	query := db.New(handler.DbPool)
	query.RemoveWeightMeasure(r.Context(), weightID)
	w.WriteHeader(http.StatusNoContent)
}
