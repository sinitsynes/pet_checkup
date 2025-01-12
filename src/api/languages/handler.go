package languages

import (
	"encoding/json"
	"log"
	"net/http"
	db "pet_checkup/internal/db/languages"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LanguagesHandler struct {
	DbPool *pgxpool.Pool
}

func (handler *LanguagesHandler) GetAvailableLanguages(w http.ResponseWriter, r *http.Request) {
	query := db.New(handler.DbPool)
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

func (handler *LanguagesHandler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	lang := &db.CreateLanguageParams{}
	decoderErr := json.NewDecoder(r.Body).Decode(lang)

	if decoderErr != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if lang.Label == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing Label"))
		return
	}
	query := db.New(handler.DbPool)
	query.CreateLanguage(r.Context(), *lang)
	w.WriteHeader(http.StatusCreated)
}
