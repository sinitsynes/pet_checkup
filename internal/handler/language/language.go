package language

import (
	"log/slog"
	"net/http"
	"pet_checkup/internal/common"
	"pet_checkup/internal/repository/language"
)

type LanguageHandler struct {
	Queries *language.Queries
}

func NewLanguageHandler(queries *language.Queries) *LanguageHandler {
	return &LanguageHandler{Queries: queries}
}

func (handler *LanguageHandler) Create() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			newLang, decodeErr := common.Decode[language.CreateParams](r)
			if decodeErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}
			created, createErr := handler.Queries.Create(r.Context(), newLang)
			if createErr != nil {
				http.Error(w, createErr.Error(), http.StatusBadRequest)
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

func (handler *LanguageHandler) GetByID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, err := common.ExtractID(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			language, repositoryErr := handler.Queries.GetByID(r.Context(), id)
			if repositoryErr != nil {
				http.Error(w, repositoryErr.Error(), http.StatusNotFound)
				return
			}
			err = common.Encode(w, r, http.StatusCreated, language)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (handler *LanguageHandler) GetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		languages, repositoryErr := handler.Queries.GetAll(r.Context())
		if repositoryErr != nil {
			http.Error(w, repositoryErr.Error(), http.StatusNotFound)
			return
		}
		err := common.Encode(w, r, http.StatusOK, languages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error("Encoding Err", "err", err.Error())
		}
	})
}

func (handler *LanguageHandler) DeleteByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := common.ExtractID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repoErr := handler.Queries.DeleteByID(r.Context(), id)
		if repoErr != nil {
			http.Error(w, repoErr.Error(), http.StatusInternalServerError)
			slog.Error("Remove Language Err", "err", repoErr.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
