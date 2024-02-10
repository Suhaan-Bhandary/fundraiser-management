package fundraiser

import (
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

// struct containing all the dependencies need for the methods
type service struct {
	fundRepo repository.FundraiserStorer
}

type Service interface {
	CreateFundraiser(userDetail dto.CreateFundraiserRequest) (int, error)
	DeleteFundraiser(fundraiserId int, tokenData dto.Token) error
	GetFundraiserDetail(fundraiserId int) (dto.FundraiserView, error)
	CloseFundraiser(fundraiserId int, tokenData dto.Token) error
	BanFundraiser(fundraiserId int) error
	ListFundraisers() ([]dto.FundraiserView, error)
	UpdateFundraiser(updateDetail dto.UpdateFundraiserRequest) error
}

func NewService(fundRepo repository.FundraiserStorer) Service {
	return &service{
		fundRepo: fundRepo,
	}
}

func (fundSvc *service) CreateFundraiser(fundDetail dto.CreateFundraiserRequest) (int, error) {
	fundraiserId, err := fundSvc.fundRepo.CreateFundraiser(fundDetail)
	if err != nil {
		return -1, err
	}

	return fundraiserId, nil
}

func (fundSvc *service) DeleteFundraiser(fundraiserId int, tokenData dto.Token) error {
	// if the role is organizer then we have to match the id and the organizer id of the fundraiser
	if tokenData.Role == constants.ORGANIZER {
		fundraiserOrganizerId, err := fundSvc.fundRepo.GetFundraiserOrganizerId(fundraiserId)
		if err != nil {
			return err
		}

		if fundraiserOrganizerId != tokenData.ID {
			return internal_errors.InvalidCredentialError{
				Message: "Only creator of fundraiser or admin can delete the fundraiser.",
			}
		}
	}

	err := fundSvc.fundRepo.DeleteFundraiser(fundraiserId)
	if err != nil {
		return err
	}

	return nil
}

func (fundSvc *service) GetFundraiserDetail(fundraiserId int) (dto.FundraiserView, error) {
	fundraiserDetail, err := fundSvc.fundRepo.GetFundraiser(fundraiserId)
	if err != nil {
		return dto.FundraiserView{}, err
	}

	return fundraiserDetail, nil
}

func (fundSvc *service) CloseFundraiser(fundraiserId int, tokenData dto.Token) error {
	// if the role is organizer then we have to match the id and the organizer id of the fundraiser
	fundraiserOrganizerId, fundraiserStatus, err := fundSvc.fundRepo.GetFundraiserOrganizerIdAndStatus(fundraiserId)
	if err != nil {
		return err
	}

	if fundraiserOrganizerId != tokenData.ID {
		return internal_errors.InvalidCredentialError{
			Message: "Only creator of fundraiser can close it",
		}
	}

	if fundraiserStatus == constants.INACTIVE_STATUS {
		return internal_errors.BadRequest{
			Message: "Fundraiser is already closed",
		}
	}

	if fundraiserStatus == constants.BANNED_STATUS {
		return internal_errors.BadRequest{
			Message: "Cannot close a banned fundraiser",
		}
	}

	err = fundSvc.fundRepo.CloseFundraiser(fundraiserId)
	if err != nil {
		return err
	}

	return nil
}

func (fundSvc *service) BanFundraiser(fundraiserId int) error {
	err := fundSvc.fundRepo.BanFundraiser(fundraiserId)
	if err != nil {
		return err
	}

	return nil
}

func (fundSvc *service) ListFundraisers() ([]dto.FundraiserView, error) {
	fundraisers, err := fundSvc.fundRepo.ListFundraiser()
	if err != nil {
		return []dto.FundraiserView{}, err
	}

	return fundraisers, nil
}

func (fundSvc *service) UpdateFundraiser(updateDetail dto.UpdateFundraiserRequest) error {
	fundraiserOrganizerId, err := fundSvc.fundRepo.GetFundraiserOrganizerId(int(updateDetail.FundraiserId))
	if err != nil {
		return err
	}

	if fundraiserOrganizerId != int(updateDetail.RequestOrganizerId) {
		return internal_errors.InvalidCredentialError{
			Message: "Only creator can update the fundraiser.",
		}
	}

	err = fundSvc.fundRepo.UpdateFundraiser(updateDetail)
	if err != nil {
		return err
	}

	return nil
}
