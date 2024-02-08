package organizer

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
	organizerRepo repository.OrganizerStorer
}

type Service interface {
	RegisterOrganizer(userDetail dto.RegisterOrganizerRequest) error
	LoginOrganizer(req dto.LoginOrganizerRequest) (string, error)
}

func NewService(organizerRepo repository.OrganizerStorer) Service {
	return &service{
		organizerRepo: organizerRepo,
	}
}

func (orgSvc *service) RegisterOrganizer(orgDetail dto.RegisterOrganizerRequest) error {
	// Hash the password before registering the user
	hashedPassword, err := helpers.HashPassword(orgDetail.Password)
	if err != nil {
		return errors.New("Internal Server Error")
	}
	orgDetail.Password = hashedPassword

	err = orgSvc.organizerRepo.RegisterOrganizer(orgDetail)
	if err != nil {
		return err
	}

	return nil
}

func (orgSvc *service) LoginOrganizer(req dto.LoginOrganizerRequest) (string, error) {
	org_id, hashedPassword, err := orgSvc.organizerRepo.GetOrganizerIDPassword(req.Email)
	if err != nil {
		return "", err
	}

	isMatch := helpers.MatchPasswordAndHash(req.Password, hashedPassword)
	if !isMatch {
		return "", internal_errors.NotFoundError{Message: "incorrect email or password"}
	}

	token, err := helpers.CreateToken(org_id, constants.ORGANIZER)
	if err != nil {
		return "", err
	}

	return token, nil
}
