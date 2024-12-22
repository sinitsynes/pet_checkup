package utils

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const PetIDParam string = "petID"

func IDfromUrl(r *http.Request, IDParamName string) int32 {
	param := chi.URLParam(r, IDParamName)
	id, int_parse_err := strconv.ParseInt(param, 10, 32)
	if int_parse_err != nil {
		log.Fatalf("tragedy occured: %q", int_parse_err)
	}
	return int32(id)
}
