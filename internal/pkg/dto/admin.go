package dto

import "errors"

type RegisterAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *RegisterAdminRequest) Validate() error {
	if req.Username == "" {
		return errors.New("username is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

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
