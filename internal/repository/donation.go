package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type DonationStorer interface {
	CreateDonation(donationDetail dto.CreateDonationRequest) (uint, error)
	ListUserDonations(userId uint) ([]dto.DonationView, error)
	ListFundraiserDonations(fundraiserId uint) ([]dto.FundraiserDonationView, error)
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
