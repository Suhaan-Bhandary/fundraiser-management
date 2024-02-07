package dto

import (
	"errors"
	"regexp"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
)

type RegisterUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (req *RegisterUserRequest) Validate() error {
	if req.FirstName == "" {
		return errors.New("first name is required")
	}

	if req.LastName == "" {
		return errors.New("last name is required")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	// check if email is in correct format
	match, err := regexp.MatchString(constants.REGEX, req.Email)
	if err != nil || !match {
		return errors.New("email is invalid")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
