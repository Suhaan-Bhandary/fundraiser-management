package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type DonationStorer interface {
	CreateDonation(donationDetail dto.CreateDonationRequest) (uint, error)
	ListUserDonations(req dto.ListUserDonationsRequest) ([]dto.DonationView, error)
	GetListUserDonationsCount(req dto.ListUserDonationsRequest) (uint, error)
	ListFundraiserDonations(fundraiserId uint) ([]dto.FundraiserDonationView, error)
	ListDonations(req dto.ListDonationsRequest) ([]dto.FundraiserDonationView, error)
	GetListDonationsCount(req dto.ListDonationsRequest) (uint, error)
}

type Donation struct {
	ID           uint
	UserId       uint
	FundraiserId uint
	Amount       float64
	IsAnonymous  bool
	CreatedAt    time.Time
}
