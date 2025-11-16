package handler

import (
	"log/slog"
	"net/http"
	"pet_checkup/internal/common"
	"pet_checkup/internal/repository/language"
	"strconv"
)

type LanguageHandler struct {
	Queries *language.Queries
}

func NewLanguageHandler(queries *language.Queries) *LanguageHandler {
	return &LanguageHandler{Queries: queries}
}

func (handler *LanguageHandler) CreateLanguage() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			newLang, decodeErr := common.Decode[language.CreateLanguageParams](r)
			if decodeErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}
			created, createErr := handler.Queries.CreateLanguage(r.Context(), newLang)
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

func (handler *LanguageHandler) GetLanguagebyID() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			langID := r.PathValue("id")
			intLangID, conversionErr := strconv.ParseInt(langID, 10, 32)
			if conversionErr != nil {
				http.Error(w, conversionErr.Error(), http.StatusBadRequest)
				return
			}
			language, repositoryErr := handler.Queries.GetLanguagebyID(r.Context(), int32(intLangID))
			if repositoryErr != nil {
				http.Error(w, repositoryErr.Error(), http.StatusNotFound)
				return
			}
			err := common.Encode(w, r, http.StatusCreated, language)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("Encoding Err", "err", err.Error())
				return
			}
		})
}

func (handler *LanguageHandler) GetLanguages() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		languages, repositoryErr := handler.Queries.GetLanguages(r.Context())
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

func (handler *LanguageHandler) RemoveLanguageByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		langID := r.PathValue("id")
		intLangID, conversionErr := strconv.ParseInt(langID, 10, 32)
		if conversionErr != nil {
			http.Error(w, conversionErr.Error(), http.StatusBadRequest)
			return
		}
		repoErr := handler.Queries.RemoveLanguage(r.Context(), int32(intLangID))
		if repoErr != nil {
			http.Error(w, repoErr.Error(), http.StatusInternalServerError)
			slog.Error("Remove Language Err", "err", repoErr.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
