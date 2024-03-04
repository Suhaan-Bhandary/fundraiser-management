package api

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/gorilla/mux"
)

func decodeId(r *http.Request) (uint, error) {
	id_str := mux.Vars(r)["id"]

	id, err := strconv.Atoi(id_str)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid Id in path")
	}

	return uint(id), nil
}

func decodeTokenFromContext(ctx context.Context) (dto.Token, error) {
	if ctx == nil {
		return dto.Token{}, internal_errors.NotFoundError{Message: "Data not found"}
	}

	tokenData, ok := ctx.Value(constants.TokenKey).(dto.Token)
	if !ok {
		return dto.Token{}, internal_errors.NotFoundError{Message: "Data not found"}
	}

	return tokenData, nil
}
