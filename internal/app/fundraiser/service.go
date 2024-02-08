package fundraiser

import (
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

// struct containing all the dependencies need for the methods
type service struct {
	fundRepo repository.FundraiserStorer
}

type Service interface {
	CreateFundraiser(userDetail dto.CreateFundraiserRequest) error
}

func NewService(fundRepo repository.FundraiserStorer) Service {
	return &service{
		fundRepo: fundRepo,
	}
}

func (fundSvc *service) CreateFundraiser(fundDetail dto.CreateFundraiserRequest) error {
	err := fundSvc.fundRepo.CreateFundraiser(fundDetail)
	if err != nil {
		return err
	}

	return nil

}
