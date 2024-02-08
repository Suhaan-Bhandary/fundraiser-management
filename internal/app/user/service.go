package user

import (
	"errors"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

// struct containing all the dependencies need for the methods
type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	RegisterUser(userDetail dto.RegisterUserRequest) error
	LoginUser(req dto.LoginUserRequest) (string, error)
	GetUserList() ([]dto.UserView, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (userSvc *service) RegisterUser(userDetail dto.RegisterUserRequest) error {
	// Hash the password before registering the user
	hashedPassword, err := helpers.HashPassword(userDetail.Password)
	if err != nil {
		return errors.New("Internal Server Error")
	}
	userDetail.Password = hashedPassword

	err = userSvc.userRepo.RegisterUser(userDetail)
	if err != nil {
		return err
	}

	return nil
}

func (userSvc *service) LoginUser(req dto.LoginUserRequest) (string, error) {
	user_id, hashedPassword, err := userSvc.userRepo.GetUserIDPassword(req.Email)
	if err != nil {
		return "", err
	}

	isMatch := helpers.MatchPasswordAndHash(req.Password, hashedPassword)
	if !isMatch {
		return "", internal_errors.NotFoundError{Message: "incorrect email or password"}
	}

	token, err := helpers.CreateToken(user_id, constants.USER)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (userSvc *service) GetUserList() ([]dto.UserView, error) {
	users, err := userSvc.userRepo.GetUserList()

	if err != nil {
		return []dto.UserView{}, err
	}

	return users, nil
}
