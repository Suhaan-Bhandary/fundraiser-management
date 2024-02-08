package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func decodeId(r *http.Request) (int, error) {
	id_str := mux.Vars(r)["id"]

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return -1, errors.New("Invalid Id in path")
	}

	return id, nil
}
