package language

import (
	"encoding/json"
	"log"
	"net/http"
	"pet_checkup/internal/db/languages"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LanguageHandler struct {
	DbPool *pgxpool.Pool
}

func (handler *LanguageHandler) GetAvailableLanguages(w http.ResponseWriter, r *http.Request) {
	query := languages.New(handler.DbPool)
	langs, err := query.GetLanguages(r.Context())
	if err != nil {
		log.Fatal(err.Error())
	}
	res, jsonErr := json.Marshal(langs)
	if jsonErr != nil {
		log.Fatal(err.Error())
	}
	w.Write(res)
}

func (handler *LanguageHandler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	lang := &languages.CreateLanguageParams{}
	decoderErr := json.NewDecoder(r.Body).Decode(lang)

	if decoderErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if lang.Label == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing Label"))
		return
	}
	query := languages.New(handler.DbPool)
	query.CreateLanguage(r.Context(), *lang)
	w.WriteHeader(http.StatusCreated)
}
