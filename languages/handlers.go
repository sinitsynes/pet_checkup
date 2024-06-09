package languages

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/go-chi/chi"
)

type Language struct {
	Name string `json:"name"`
}

type LanguageService struct {
	http.Handler
}

func (s *LanguageService) GetAvailableLanguages(w http.ResponseWriter, r *http.Request) {
	langs, err := GetLanguages(r.Context())
	if err != nil {
		log.Fatalf(err.Error())
	}
	list := []Language{}
	for _, lang_name := range langs {
		list = append(list, Language{lang_name})
	}
	res, _ := json.Marshal(list)
	w.Write(res)
}

func LanguageRouter() *LanguageService {
	s := new(LanguageService)
	router := chi.NewRouter()
	router.Get("/", s.GetAvailableLanguages)
	s.Handler = router
	return s
}
