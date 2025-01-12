package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func IDfromUrl(r *http.Request, IDParamName string) (int32, error) {
	param := chi.URLParam(r, IDParamName)
	id, int_parse_err := strconv.ParseInt(param, 10, 32)
	if int_parse_err != nil {
		return int32(id), errors.New("can't parse resource identifier")
	}
	return int32(id), nil
}
