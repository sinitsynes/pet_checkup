package pet

import (
	"log/slog"
	"net/http"
	"pet_checkup/internal/common"
	"pet_checkup/internal/repository/pet"
)

type PetHandler struct {
	Queries *pet.Queries
}

func NewPetHandler(queries *pet.Queries) *PetHandler {
	return &PetHandler{Queries: queries}
}

func (handler *PetHandler) GetAll() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			pets, repoErr := handler.Queries.GetAll(r.Context())
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			err := common.Encode(w, r, http.StatusOK, pets)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (handler *PetHandler) Create() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			newPet, decodeErr := common.Decode[pet.CreateParams](r)
			if decodeErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}
			created, repoErr := handler.Queries.Create(r.Context(), newPet)
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			err := common.Encode(w, r, http.StatusCreated, created)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (handler *PetHandler) GetByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := common.ExtractID(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			pet, repoErr := handler.Queries.GetByID(r.Context(), id)
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			err = common.Encode(w, r, http.StatusOK, pet)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (handler *PetHandler) UpdateByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := common.ExtractID(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatePet, decodeErr := common.Decode[pet.UpdateByIDParams](r)
			if decodeErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}
			updatePet.ID = id
			updated, repoErr := handler.Queries.UpdateByID(r.Context(), updatePet)
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

func (handler *PetHandler) DeleteByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := common.ExtractID(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			repoErr := handler.Queries.DeleteByID(r.Context(), id)
			if repoErr != nil {
				http.Error(w, repoErr.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
}
