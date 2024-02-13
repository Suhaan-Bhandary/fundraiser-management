package dto

import (
	"errors"
	"slices"
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
)

type FundraiserView struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	OrganizerId   uint      `json:"organizer_id"`
	OrganizerName string    `json:"organizer_name"`
	ImageUrl      string    `json:"image_url"`
	VideoUrl      string    `json:"video_url"`
	TargetAmount  float64   `json:"target_amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateFundraiserRequest struct {
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	OrganizerId  uint    `json:"organizer_id"`
	ImageUrl     string  `json:"image_url"`
	VideoUrl     string  `json:"video_url"`
	TargetAmount float64 `json:"target_amount"`
}

func (req *CreateFundraiserRequest) Validate() error {
	if req.Title == "" {
		return errors.New("title is required")
	}

	if req.Description == "" {
		return errors.New("description is required")
	}

	if req.OrganizerId == 0 {
		return errors.New("organizer id is required")
	}

	if req.ImageUrl == "" {
		return errors.New("image URL is required")
	}

	if req.VideoUrl == "" {
		return errors.New("video URL is required")
	}

	if req.TargetAmount <= float64(0) {
		return errors.New("target amount is required")
	}

	return nil
}

type DeleteFundraiserRequest struct {
	Token
	FundraiserId uint `json:"fundraiser_id"`
}

type CreateFundraiserResponse struct {
	FundraiserId uint `json:"fundraiser_id"`
}

type GetFundraiserResponse struct {
	Fundraiser FundraiserView `json:"fundraiser"`
}

type ListFundraisersResponse struct {
	Fundraisers []FundraiserView `json:"fundraisers"`
	TotalCount  uint             `json:"total_count"`
}

type ListFundraiserDonationsResponse struct {
	Donations  []FundraiserDonationView `json:"donations"`
	TotalCount uint                     `json:"total_count"`
}

type ListDonationsResponse struct {
	Donations  []FundraiserDonationView `json:"donations"`
	TotalCount uint                     `json:"total_count"`
}

type UpdateFundraiserRequest struct {
	RequestOrganizerId uint    `json:"request_organizer_id"`
	FundraiserId       uint    `json:"fundraiser_id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	ImageUrl           string  `json:"image_url"`
	VideoUrl           string  `json:"video_url"`
	TargetAmount       float64 `json:"target_amount"`
}

func (req *UpdateFundraiserRequest) Validate() error {
	if req.RequestOrganizerId <= 0 {
		return errors.New("invalid organizer id in token")
	}

	if req.FundraiserId <= 0 {
		return errors.New("invalid fundraiser id")
	}

	if req.Title == "" {
		return errors.New("title is required")
	}

	if req.Description == "" {
		return errors.New("description is required")
	}

	if req.ImageUrl == "" {
		return errors.New("image URL is required")
	}

	if req.VideoUrl == "" {
		return errors.New("video URL is required")
	}

	if req.TargetAmount <= float64(0) {
		return errors.New("non negative target amount is required")
	}

	return nil
}

type ListFundraisersRequest struct {
	Search             string `json:"search"`
	Status             string `json:"status"`
	OrderByKey         string `json:"order_by"`
	OrderByIsAscending bool   `json:"is_ascending"`
	Offset             uint   `json:"offset"`
	Limit              uint   `json:"limit"`
}

func (req *ListFundraisersRequest) Validate() error {
	if req.Offset < 0 {
		return errors.New("invalid offset, offset cannot be negative")
	}

	if req.Limit <= 0 {
		return errors.New("limit should be positive number")
	}

	if req.Limit > 1000 {
		return errors.New("Limit should be in range [1, 1000]")
	}

	allStatus := []string{constants.ACTIVE_STATUS, constants.INACTIVE_STATUS, constants.BANNED_STATUS}
	if req.OrderByKey != "" && !slices.Contains(allStatus, req.Status) {
		return errors.New("invalid status key")
	}

	orderKeys := []string{
		"fundraiser_id", "organizer_id", "title", "description", "organizer_name",
		"target_amount", "status", "created_at", "updated_at",
	}

	if req.OrderByKey != "" && !slices.Contains(orderKeys, req.OrderByKey) {
		return errors.New("invalid order key")
	}

	return nil
}
