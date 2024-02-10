package dto

import (
	"errors"
	"time"
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

type CreateFundraiserResponse struct {
	FundraiserId int `json:"fundraiser_id"`
}

type GetFundraiserResponse struct {
	Fundraiser FundraiserView `json:"fundraiser"`
}

type ListFundraisersResponse struct {
	Fundraisers []FundraiserView `json:"fundraisers"`
}

type ListFundraiserDonationsResponse struct {
	Donations []FundraiserDonationView `json:"donations"`
}

type ListDonationsResponse struct {
	Donations []FundraiserDonationView `json:"donations"`
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
