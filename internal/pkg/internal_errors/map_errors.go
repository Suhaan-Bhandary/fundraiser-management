package internal_errors

import "net/http"

func MatchError(err error) (int, error) {
	switch err.(type) {
	case DuplicateKeyError:
		return http.StatusBadRequest, err
	case InvalidCredentialError:
		return http.StatusUnauthorized, err
	case NotFoundError:
		return http.StatusNotFound, err
	case BadRequest:
		return http.StatusBadRequest, err
	default:
		return http.StatusInternalServerError, err
	}
}
