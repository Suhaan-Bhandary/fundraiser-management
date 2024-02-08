package dto

import (
	"errors"
	"regexp"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
)

type UserView struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
}

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

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginUserRequest) Validate() error {
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

type LoginUserResponse struct {
	Token string `json:"token"`
}

type GetUsersResponse struct {
	Users []UserView `json:"users"`
}
