package pets

import (
	"encoding/json"
	"net/http"
	"pet_checkup/api/utils"
	db "pet_checkup/internal/db/pets"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PetHandler struct {
	DbPool *pgxpool.Pool
}

func (handler *PetHandler) CreateRecord(w http.ResponseWriter, r *http.Request) {
	createPetParams := &db.CreatePetParams{}
	decoderErr := json.NewDecoder(r.Body).Decode(createPetParams)
	if decoderErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := db.New(handler.DbPool)
	if dbErr := query.CreatePet(r.Context(), *createPetParams); dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *PetHandler) EditRecord(w http.ResponseWriter, r *http.Request) {
	petID, err := utils.IDfromUrl(r, "PetID")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	updatePetParams := &db.UpdatePetParams{ID: petID}
	decoder_err := json.NewDecoder(r.Body).Decode(updatePetParams)
	if decoder_err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := db.New(handler.DbPool)
	if db_err := query.UpdatePet(r.Context(), *updatePetParams); db_err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (handler *PetHandler) GetRecord(w http.ResponseWriter, r *http.Request) {
	petID, err := utils.IDfromUrl(r, "PetIDParam")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	query := db.New(handler.DbPool)
	pet, db_err := query.GetPet(r.Context(), petID)
	if db_err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (handler *PetHandler) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	petID, err := utils.IDfromUrl(r, "PetIDParam")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	query := db.New(handler.DbPool)
	if err := query.DeletePet(r.Context(), petID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
