package language

import (
	"net/http"
	"pet_checkup/internal/handler"
	"pet_checkup/internal/repository/language"

	"github.com/jackc/pgx/v5/pgxpool"
)

func makeLanguageRouter(languageHandler *handler.LanguageHandler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /languages", languageHandler.GetLanguages())
	mux.Handle("POST /languages", languageHandler.CreateLanguage())
	mux.Handle("GET /languages/{id}", languageHandler.GetLanguagebyID())

	return mux
}

func LanguageRouter(DBPool *pgxpool.Pool) http.Handler {
	queries := language.New(DBPool)
	handler := handler.NewLanguageHandler(queries)
	return makeLanguageRouter(handler)
}
