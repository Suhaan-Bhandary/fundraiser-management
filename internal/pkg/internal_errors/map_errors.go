package internal_errors

import "net/http"

func MatchError(err error) (int, error) {
	switch err.(type) {
	case DuplicateKeyError:
		return http.StatusBadRequest, err
	default:
		return http.StatusInternalServerError, err
	}
}
