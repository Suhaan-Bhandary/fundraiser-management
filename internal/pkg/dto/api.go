package dto

import "errors"

type IdParam struct {
	ID string
}

func (data *IdParam) validate() error {
	if data.ID == "" {
		return errors.New("Id is required")
	}

	return nil
}

type MessageResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
