package api

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
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

func decodeTokenFromContext(ctx context.Context) (dto.Token, error) {
	if ctx == nil {
		return dto.Token{}, internal_errors.NotFoundError{Message: "Data not found"}
	}

	tokenData, ok := ctx.Value("token-data").(dto.Token)
	if !ok {
		return dto.Token{}, internal_errors.NotFoundError{Message: "Data not found"}
	}

	return tokenData, nil
}
