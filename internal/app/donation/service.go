package donation

import (
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

type service struct {
	donationRepo   repository.DonationStorer
	fundraiserRepo repository.FundraiserStorer
}

type Service interface {
	CreateDonation(donationDetail dto.CreateDonationRequest) (uint, error)
	ListUserDonation(req dto.ListUserDonationsRequest) ([]dto.DonationView, uint, error)
	ListFundraiserDonations(req dto.ListFundraiserDonationsRequest) ([]dto.FundraiserDonationView, uint, error)
	ListDonations(req dto.ListDonationsRequest) ([]dto.FundraiserDonationView, uint, error)
}

func NewService(donationRepo repository.DonationStorer, fundraiserRepo repository.FundraiserStorer) Service {
	return &service{
		donationRepo:   donationRepo,
		fundraiserRepo: fundraiserRepo,
	}
}

func (donationSvc *service) CreateDonation(donationDetail dto.CreateDonationRequest) (uint, error) {
	fundraiserStatus, err := donationSvc.fundraiserRepo.GetFundraiserStatus(donationDetail.FundraiserId)
	if err != nil {
		return 0, err
	}

	if fundraiserStatus != constants.ACTIVE_STATUS {
		return 0, internal_errors.BadRequest{Message: "cannot donate as fundraiser is not active"}
	}

	donationId, err := donationSvc.donationRepo.CreateDonation(donationDetail)
	if err != nil {
		return 0, err
	}

	// Update Fundraiser status according to total amount collected and target amount
	err = donationSvc.fundraiserRepo.UpdateFundraiserStatus(donationDetail.FundraiserId)

	return donationId, nil
}

func (donationSvc *service) ListUserDonation(req dto.ListUserDonationsRequest) ([]dto.DonationView, uint, error) {
	totalCount, err := donationSvc.donationRepo.GetListUserDonationsCount(req)
	if err != nil {
		return []dto.DonationView{}, 0, err
	}

	userDonations, err := donationSvc.donationRepo.ListUserDonations(req)
	if err != nil {
		return []dto.DonationView{}, 0, err
	}

	return userDonations, totalCount, nil
}

func (donationSvc *service) ListFundraiserDonations(req dto.ListFundraiserDonationsRequest) ([]dto.FundraiserDonationView, uint, error) {
	totalCount, err := donationSvc.donationRepo.GetListFundraiserDonationsCount(req)
	if err != nil {
		return []dto.FundraiserDonationView{}, 0, err
	}

	fundraiserDonations, err := donationSvc.donationRepo.ListFundraiserDonations(req)
	if err != nil {
		return []dto.FundraiserDonationView{}, 0, err
	}

	return fundraiserDonations, totalCount, nil
}

func (donationSvc *service) ListDonations(req dto.ListDonationsRequest) ([]dto.FundraiserDonationView, uint, error) {
	totalCount, err := donationSvc.donationRepo.GetListDonationsCount(req)
	if err != nil {
		return []dto.FundraiserDonationView{}, 0, err
	}

	donations, err := donationSvc.donationRepo.ListDonations(req)
	if err != nil {
		return []dto.FundraiserDonationView{}, 0, err
	}

	return donations, totalCount, nil
}
