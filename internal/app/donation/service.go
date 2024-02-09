package donation

import (
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

type service struct {
	donationRepo repository.DonationStorer
}

type Service interface {
	CreateDonation(donationDetail dto.CreateDonationRequest) (int, error)
}

func NewService(donationRepo repository.DonationStorer) Service {
	return &service{
		donationRepo: donationRepo,
	}
}

func (donationSvc *service) CreateDonation(donationDetail dto.CreateDonationRequest) (int, error) {
	donationId, err := donationSvc.donationRepo.CreateDonation(donationDetail)
	if err != nil {
		return -1, err
	}

	return donationId, nil
}
