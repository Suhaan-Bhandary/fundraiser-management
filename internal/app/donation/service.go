package donation

import (
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

type service struct {
	donationRepo repository.DonationStorer
}

type Service interface {
	CreateDonation(donationDetail dto.CreateDonationRequest) (uint, error)
	ListUserDonation(userId uint) ([]dto.DonationView, error)
	ListFundraiserDonations(fundraiserId uint) ([]dto.FundraiserDonationView, error)
	ListDonations() ([]dto.FundraiserDonationView, error)
}

func NewService(donationRepo repository.DonationStorer) Service {
	return &service{
		donationRepo: donationRepo,
	}
}

func (donationSvc *service) CreateDonation(donationDetail dto.CreateDonationRequest) (uint, error) {
	donationId, err := donationSvc.donationRepo.CreateDonation(donationDetail)
	if err != nil {
		return 0, err
	}

	return donationId, nil
}

func (donationSvc *service) ListUserDonation(userId uint) ([]dto.DonationView, error) {
	userDonations, err := donationSvc.donationRepo.ListUserDonations(userId)
	if err != nil {
		return []dto.DonationView{}, err
	}

	return userDonations, nil
}

func (donationSvc *service) ListFundraiserDonations(fundraiserId uint) ([]dto.FundraiserDonationView, error) {
	fundraiserDonations, err := donationSvc.donationRepo.ListFundraiserDonations(fundraiserId)
	if err != nil {
		return []dto.FundraiserDonationView{}, err
	}

	return fundraiserDonations, nil
}

func (donationSvc *service) ListDonations() ([]dto.FundraiserDonationView, error) {
	donations, err := donationSvc.donationRepo.ListDonations()
	if err != nil {
		return []dto.FundraiserDonationView{}, err
	}

	return donations, nil
}
