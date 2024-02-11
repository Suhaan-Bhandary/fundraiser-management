package dto

import (
	"errors"
	"slices"
	"time"
)

type DonationView struct {
	ID              uint      `json:"id"`
	FundraiserId    uint      `json:"fundraiser_id"`
	FundraiserTitle string    `json:"fundraiser_title"`
	Amount          float64   `json:"amount"`
	IsAnonymous     string    `json:"is_anonymous"`
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
	DonationId uint `json:"donation_id"`
}

type ListUserDonationsResponse struct {
	Donations  []DonationView `json:"donations"`
	TotalCount uint           `json:"total_count"`
}

type ListUserDonationsRequest struct {
	UserId             uint   `json:"user_id"`
	Search             string `json:"search"`
	IsAnonymous        string `json:"is_anonymous"`
	Offset             uint   `json:"offset"`
	Limit              uint   `json:"limit"`
	OrderByKey         string `json:"order_by"`
	OrderByIsAscending bool   `json:"is_ascending"`
}

func (req *ListUserDonationsRequest) Validate() error {
	if req.Offset < 0 {
		return errors.New("invalid offset, offset cannot be negative")
	}

	if req.Limit <= 0 {
		return errors.New("limit should be positive number")
	}

	if req.Limit > 1000 {
		return errors.New("Limit should be in range [1, 1000]")
	}

	orderKeys := []string{"id", "fundraiser_title", "amount", "created_at", "updated_at"}
	if req.OrderByKey != "" && !slices.Contains(orderKeys, req.OrderByKey) {
		return errors.New("invalid order key")
	}

	return nil
}
