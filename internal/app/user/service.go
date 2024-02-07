package user

import (
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

// struct containing all the dependencies need for the methods
type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	RegisterUser(userDetail dto.RegisterUserRequest) error
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (orderSvc *service) RegisterUser(userDetail dto.RegisterUserRequest) error {
	err := orderSvc.userRepo.RegisterUser(userDetail)
	if err != nil {
		return err
	}

	return nil
}
