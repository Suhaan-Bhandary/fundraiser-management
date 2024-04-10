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
	CreateFundraiser(userDetail dto.CreateFundraiserRequest) (uint, error)
	DeleteFundraiser(req dto.DeleteFundraiserRequest) error
	GetFundraiserDetail(fundraiserId uint) (dto.FundraiserView, error)
	CloseFundraiser(fundraiserId uint, tokenData dto.Token) error
	BanFundraiser(fundraiserId uint) error
	UnBanFundraiser(fundraiserId uint) error
	ListFundraisers(req dto.ListFundraisersRequest) ([]dto.FundraiserView, uint, error)
	UpdateFundraiser(updateDetail dto.UpdateFundraiserRequest) error
}

func NewService(fundRepo repository.FundraiserStorer) Service {
	return &service{
		fundRepo: fundRepo,
	}
}

func (fundSvc *service) CreateFundraiser(fundDetail dto.CreateFundraiserRequest) (uint, error) {
	fundraiserId, err := fundSvc.fundRepo.CreateFundraiser(fundDetail)
	if err != nil {
		return 0, err
	}

	return fundraiserId, nil
}

func (fundSvc *service) DeleteFundraiser(req dto.DeleteFundraiserRequest) error {
	// if the role is organizer then we have to match the id and the organizer id of the fundraiser
	if req.Token.Role == constants.ORGANIZER {
		fundraiserOrganizerId, err := fundSvc.fundRepo.GetFundraiserOrganizerId(req.FundraiserId)
		if err != nil {
			return err
		}

		if fundraiserOrganizerId != req.Token.ID {
			return internal_errors.InvalidCredentialError{
				Message: "Only creator of fundraiser or admin can delete the fundraiser.",
			}
		}
	}

	err := fundSvc.fundRepo.DeleteFundraiser(req.FundraiserId)
	if err != nil {
		return err
	}

	return nil
}

func (fundSvc *service) GetFundraiserDetail(fundraiserId uint) (dto.FundraiserView, error) {
	fundraiserDetail, err := fundSvc.fundRepo.GetFundraiser(fundraiserId)
	if err != nil {
		return dto.FundraiserView{}, err
	}

	return fundraiserDetail, nil
}

func (fundSvc *service) CloseFundraiser(fundraiserId uint, tokenData dto.Token) error {
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

func (fundSvc *service) BanFundraiser(fundraiserId uint) error {
	err := fundSvc.fundRepo.BanFundraiser(fundraiserId)
	if err != nil {
		return err
	}

	return nil
}

func (fundSvc *service) UnBanFundraiser(fundraiserId uint) error {
	err := fundSvc.fundRepo.UnBanFundraiser(fundraiserId)
	if err != nil {
		return err
	}

	return nil
}

func (fundSvc *service) ListFundraisers(req dto.ListFundraisersRequest) ([]dto.FundraiserView, uint, error) {
	totalCount, err := fundSvc.fundRepo.GetListFundraisersCount(req)
	if err != nil {
		return []dto.FundraiserView{}, 0, err
	}

	fundraisers, err := fundSvc.fundRepo.ListFundraiser(req)
	if err != nil {
		return []dto.FundraiserView{}, 0, err
	}

	return fundraisers, totalCount, nil
}

func (fundSvc *service) UpdateFundraiser(updateDetail dto.UpdateFundraiserRequest) error {
	fundraiserOrganizerId, err := fundSvc.fundRepo.GetFundraiserOrganizerId(updateDetail.FundraiserId)
	if err != nil {
		return err
	}

	if fundraiserOrganizerId != updateDetail.RequestOrganizerId {
		return internal_errors.InvalidCredentialError{
			Message: "Only creator can update the fundraiser.",
		}
	}

	err = fundSvc.fundRepo.UpdateFundraiser(updateDetail)
	if err != nil {
		return err
	}

	// Update Fundraiser status according to total amount collected and target amount
	err = fundSvc.fundRepo.UpdateFundraiserStatus(updateDetail.FundraiserId)

	return nil
}
