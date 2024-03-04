package admin

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
	adminRepo repository.AdminStorer
}

type Service interface {
	RegisterAdmin(req dto.RegisterAdminRequest) error
	LoginAdmin(req dto.LoginAdminRequest) (string, error)
}

func NewService(adminRepo repository.AdminStorer) Service {
	return &service{
		adminRepo: adminRepo,
	}
}

func (adminSvc *service) RegisterAdmin(req dto.RegisterAdminRequest) error {
	// Hash the password before registering the user
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return errors.New("Internal Server Error")
	}
	req.Password = hashedPassword

	err = adminSvc.adminRepo.RegisterAdmin(req)
	if err != nil {
		return err
	}

	return nil
}

func (adminSvc *service) LoginAdmin(req dto.LoginAdminRequest) (string, error) {
	userId, hashedPassword, err := adminSvc.adminRepo.GetAdminIDPassword(req.Username)
	if err != nil {
		return "", err
	}

	isMatch := helpers.MatchPasswordAndHash(req.Password, hashedPassword)
	if !isMatch {
		return "", internal_errors.NotFoundError{Message: "incorrect email or password"}
	}

	token, err := helpers.CreateToken(userId, constants.ADMIN)
	if err != nil {
		return "", err
	}

	return token, nil
}
