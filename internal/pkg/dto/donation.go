package dto

import (
	"errors"
	"time"
)

type DonationView struct {
	ID              uint      `json:"id"`
	FundraiserId    uint      `json:"fundraiser_id"`
	FundraiserTitle string    `json:"fundraiser_title"`
	Amount          float64   `json:"amount"`
	IsAnonymous     bool      `json:"is_anonymous"`
	CreatedAt       time.Time `json:"created_at"`
}

type FundraiserDonationView struct {
	ID              uint      `json:"id"`
	UserId          uint      `json:"user_id"`
	UserName        string    `json:"user_name"`
	FundraiserId    uint      `json:"fundraiser_id"`
	FundraiserTitle string    `json:"fundraiser_title"`
	Amount          float64   `json:"amount"`
	IsAnonymous     bool      `json:"is_anonymous"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateDonationRequest struct {
	UserId       uint    `json:"user_id"`
	FundraiserId uint    `json:"fundraiser_id"`
	Amount       float64 `json:"amount"`
	IsAnonymous  bool    `json:"is_anonymous"`
}

func (req *CreateDonationRequest) Validate() error {
	if req.UserId == 0 {
		return errors.New("user id is required")
	}

	if req.FundraiserId == 0 {
		return errors.New("fundraiser id is required")
	}

	if req.Amount <= float64(0) {
		return errors.New("amount required and should be non negative")
	}

	return nil
}

type CreateDonationResponse struct {
	DonationId int `json:"donation_id"`
}

type ListUserDonationsResponse struct {
	Donations []DonationView `json:"donations"`
}
