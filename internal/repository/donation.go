package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type DonationStorer interface {
	CreateDonation(donationDetail dto.CreateDonationRequest) (int, error)
	ListUserDonations(user_id int) ([]dto.DonationView, error)
	ListFundraiserDonations(fundraiser_id int) ([]dto.FundraiserDonationView, error)
	ListDonations() ([]dto.FundraiserDonationView, error)
}

type Donation struct {
	ID           uint
	UserId       uint
	FundraiserId uint
	Amount       float64
	IsAnonymous  bool
	CreatedAt    time.Time
}
