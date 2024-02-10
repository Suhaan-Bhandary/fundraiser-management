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
	ListUserDonation(user_id int) ([]dto.DonationView, error)
	ListFundraiserDonations(fundraiserId int) ([]dto.FundariserDonationView, error)
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

func (donationSvc *service) ListUserDonation(user_id int) ([]dto.DonationView, error) {
	userDonations, err := donationSvc.donationRepo.ListUserDonations(user_id)
	if err != nil {
		return []dto.DonationView{}, err
	}

	return userDonations, nil
}

func (donationSvc *service) ListFundraiserDonations(fundraiserId int) ([]dto.FundariserDonationView, error) {
	fundraiserDonations, err := donationSvc.donationRepo.ListFundraiserDonations(fundraiserId)
	if err != nil {
		return []dto.FundariserDonationView{}, err
	}

	return fundraiserDonations, nil
}
