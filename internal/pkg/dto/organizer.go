package dto

import (
	"errors"
	"regexp"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
)

type OrganizerView struct {
	ID           uint
	Organization string
	Detail       string
	Email        string
	Mobile       string
	IsVerified   bool
}

type RegisterOrganizerRequest struct {
	Organization string `json:"organization"`
	Detail       string `json:"detail"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Mobile       string `json:"mobile"`
}

func (req *RegisterOrganizerRequest) Validate() error {
	if req.Organization == "" {
		return errors.New("Organization name required")
	}

	if req.Detail == "" {
		return errors.New("Organization detail required")
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

	if req.Password == "" {
		return errors.New("mobile is required")
	}

	return nil
}

type LoginOrganizerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginOrganizerRequest) Validate() error {
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

type GetNotVerifiedOrganizersResponse struct {
	Organizers []OrganizerView `json:"organizers"`
}
