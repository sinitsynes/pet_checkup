package common

import (
	"net/http"
	"strconv"
)

func ExtractID(r *http.Request) (int32, error) {
	id := r.PathValue("id")
	intID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(intID), nil
}
