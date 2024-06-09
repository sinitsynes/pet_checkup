package pets

import (
	"encoding/json"
	"log"
	"net/http"
	db "pet_checkup/db/pets"
	"strconv"

	"github.com/go-chi/chi"
)

func petID(r *http.Request) int32 {
	param := chi.URLParam(r, "petID")
	id, int_parse_err := strconv.ParseInt(param, 10, 32)
	if int_parse_err != nil {
		log.Fatalf("tragedy occured: %q", int_parse_err)
	}
	return int32(id)
}

type PetService struct {
	http.Handler
}

func (s *PetService) CreateRecord(w http.ResponseWriter, r *http.Request) {
	createPetParams := &db.CreatePetParams{}
	decoderErr := json.NewDecoder(r.Body).Decode(createPetParams)
	if decoderErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if db_err := CreatePet(r.Context(), *createPetParams); db_err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *PetService) EditRecord(w http.ResponseWriter, r *http.Request) {
	editPetParams := &db.UpdatePetParams{ID: petID(r)}
	decoder_err := json.NewDecoder(r.Body).Decode(editPetParams)
	if decoder_err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if db_err := EditPet(r.Context(), *editPetParams); db_err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusAccepted)
}

func (s *PetService) GetRecord(w http.ResponseWriter, r *http.Request) {
	pet, db_err := GetPet(r.Context(), petID(r))
	if db_err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	res, err := json.Marshal(pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(res)
}

func (s *PetService) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	if err := DeletePet(r.Context(), petID(r)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *PetService) AddWeightRecord(w http.ResponseWriter, r *http.Request) {
	addWeightParams := &db.CreateWeightMeasureParams{}
	if dbErr := CreateWeightMeasure(r.Context(), *addWeightParams); dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *PetService) DeleteWeightRecord(w http.ResponseWriter, r *http.Request) {
	removeWeightParams := &db.RemoveWeightMeasureParams{}
	if decoderErr := json.NewDecoder(r.Body).Decode(removeWeightParams); decoderErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}

func PetRouter() *PetService {
	s := new(PetService)
	router := chi.NewRouter()
	router.Post("/", s.CreateRecord)
	router.Put("/{petID}", s.EditRecord)
	router.Get("/{petID}", s.GetRecord)
	router.Delete("/{petID}", s.DeleteRecord)
	s.Handler = router
	return s
}
