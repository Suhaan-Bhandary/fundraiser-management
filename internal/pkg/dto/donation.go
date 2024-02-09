package dto

import "errors"

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
