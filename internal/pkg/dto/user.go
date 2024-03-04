package dto

import (
	"errors"
	"regexp"
	"slices"

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
	match, err := regexp.MatchString(constants.EMAIL_REGEX, req.Email)
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
	match, err := regexp.MatchString(constants.EMAIL_REGEX, req.Email)
	if err != nil || !match {
		return errors.New("email is invalid")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type GetUsersResponse struct {
	Users      []UserView `json:"users"`
	TotalCount uint       `json:"total_count"`
}

type GetUserProfileResponse struct {
	User UserView `json:"user"`
}

type ListUserRequest struct {
	Search             string `json:"search"`
	Offset             uint   `json:"offset"`
	Limit              uint   `json:"limit"`
	OrderByKey         string `json:"order_by"`
	OrderByIsAscending bool   `json:"is_ascending"`
}

func (req *ListUserRequest) Validate() error {
	if req.Offset < 0 {
		return errors.New("invalid offset, offset cannot be negative")
	}

	if req.Limit <= 0 {
		return errors.New("limit should be positive number")
	}

	if req.Limit > 1000 {
		return errors.New("Limit should be in range [1, 1000]")
	}

	orderKeys := []string{"id", "first_name", "last_name", "email", "created_at", "updated_at"}
	if req.OrderByKey != "" && !slices.Contains(orderKeys, req.OrderByKey) {
		return errors.New("invalid order key")
	}

	return nil
}
