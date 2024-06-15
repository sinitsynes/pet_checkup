package panels

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type PanelService struct {
	http.Handler
}

type CreateMeasurementAPIRequest struct {
	Name      string `json:"name"`
	LangLabel string `json:"lang_label"`
}

func (s *PanelService) CreatePanelMeasurement(w http.ResponseWriter, r *http.Request) {
	apiRequest := new(CreateMeasurementAPIRequest)
	if decodeErr := json.NewDecoder(r.Body).Decode(apiRequest); decodeErr != nil {
		log.Fatalf("decoder err %q", decodeErr)
	}
	if err := AddPanelMeasurement(r.Context(), *apiRequest); err != nil {
		w.Write([]byte("cant create measurement"))
	}
}

func PanelsRouter() *PanelService {
	s := new(PanelService)
	router := chi.NewRouter()
	router.Post("/measurement", s.CreatePanelMeasurement)
	s.Handler = router
	return s
}
