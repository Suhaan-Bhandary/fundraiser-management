package dto

import (
	"errors"
	"fmt"
	"regexp"
	"slices"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
)

type OrganizerView struct {
	ID         uint
	Name       string
	Detail     string
	Email      string
	Mobile     string
	IsVerified bool
}

type RegisterOrganizerRequest struct {
	Name     string `json:"name"`
	Detail   string `json:"detail"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

func (req *RegisterOrganizerRequest) Validate() error {
	if req.Name == "" {
		return errors.New("organizer name required")
	}

	if req.Detail == "" {
		return errors.New("organizer detail required")
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

	match, err = regexp.MatchString(constants.MOBILE_REGEX, req.Mobile)
	if err != nil || !match {
		fmt.Println("#####################")
		fmt.Println(err, match)
		return errors.New("mobile is invalid")
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

type ListOrganizersResponse struct {
	Organizers []OrganizerView `json:"organizers"`
	TotalCount uint            `json:"total_count"`
}

type GetOrganizerResponse struct {
	Organizer OrganizerView `json:"organizer"`
}

type UpdateOrganizerRequest struct {
	OrganizerId uint   `json:"organizer_id"`
	Detail      string `json:"detail"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
}

func (req *UpdateOrganizerRequest) Validate() error {
	if req.OrganizerId <= 0 {
		return errors.New("invalid organizer id")
	}

	if req.Detail == "" {
		return errors.New("organizer detail required")
	}

	if req.Email == "" {
		return errors.New("email is required")
	}

	// check if email is in correct format
	match, err := regexp.MatchString(constants.EMAIL_REGEX, req.Email)
	if err != nil || !match {
		return errors.New("email is invalid")
	}

	match, err = regexp.MatchString(constants.MOBILE_REGEX, req.Mobile)
	if err != nil || !match {
		return errors.New("mobile is invalid")
	}

	return nil
}

type ListOrganizersRequest struct {
	Search             string `json:"search"`
	Verified           string `json:"verified"`
	Offset             uint   `json:"offset"`
	Limit              uint   `json:"limit"`
	OrderByKey         string `json:"order_by"`
	OrderByIsAscending bool   `json:"is_ascending"`
}

func (req *ListOrganizersRequest) Validate() error {
	if req.Verified != "" && req.Verified != "true" && req.Verified != "false" {
		return errors.New("invalid value of verified, it can only be true or false")
	}

	if req.Offset < 0 {
		return errors.New("invalid offset, offset cannot be negative")
	}

	if req.Limit <= 0 {
		return errors.New("limit should be positive number")
	}

	if req.Limit > 1000 {
		return errors.New("Limit should be in range [1, 1000]")
	}

	orderKeys := []string{"id", "name", "email", "mobile", "is_verified", "created_at", "updated_at"}
	if req.OrderByKey != "" && !slices.Contains(orderKeys, req.OrderByKey) {
		return errors.New("invalid order key")
	}

	return nil
}
