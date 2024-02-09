package dto

import (
	"errors"
)

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
		return errors.New("organization id is required")
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
