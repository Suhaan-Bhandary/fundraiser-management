package dto

import "errors"

type LoginAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *LoginAdminRequest) Validate() error {
	if req.Username == "" {
		return errors.New("username is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
