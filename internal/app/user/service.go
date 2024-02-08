package user

import (
	"errors"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/helpers"
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
	// Hash the password before registering the user
	hashedPassword, err := helpers.HashPassword(userDetail.Password)
	if err != nil {
		return errors.New("Internal Server Error")
	}
	userDetail.Password = hashedPassword

	err = orderSvc.userRepo.RegisterUser(userDetail)
	if err != nil {
		return err
	}

	return nil
}
