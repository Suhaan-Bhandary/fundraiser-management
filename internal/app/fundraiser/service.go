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
