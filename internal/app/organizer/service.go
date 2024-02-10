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
	VerifyOrganizer(organizerId uint) error
	GetOrganizerList(search string, verified string) ([]dto.OrganizerView, error)
	DeleteOrganizer(organizerId uint) error
	GetOrganizer(organizerId uint) (dto.OrganizerView, error)
	UpdateOrganizer(req dto.UpdateOrganizerRequest) error
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

func (orgSvc *service) DeleteOrganizer(organizerId uint) error {
	err := orgSvc.organizerRepo.DeleteOrganizer(organizerId)
	if err != nil {
		return err
	}

	return nil
}

func (orgSvc *service) LoginOrganizer(req dto.LoginOrganizerRequest) (string, error) {
	org_id, hashedPassword, isVerified, err := orgSvc.organizerRepo.GetOrganizerIDPasswordAndVerifyStatus(req.Email)
	if err != nil {
		return "", err
	}

	if !isVerified {
		return "", internal_errors.InvalidCredentialError{Message: "Organizer is not verified, please contact admin"}
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

func (orgSvc *service) VerifyOrganizer(organizerId uint) error {
	err := orgSvc.organizerRepo.VerifyOrganizer(organizerId)
	if err != nil {
		return err
	}

	return nil
}

func (orgSvc *service) GetOrganizerList(search string, verified string) ([]dto.OrganizerView, error) {
	organizers, err := orgSvc.organizerRepo.GetOrganizerList(search, verified)
	if err != nil {
		return []dto.OrganizerView{}, err
	}

	return organizers, nil
}

func (orgSvc *service) GetOrganizer(organizerId uint) (dto.OrganizerView, error) {
	organizer, err := orgSvc.organizerRepo.GetOrganizer(organizerId)
	if err != nil {
		return dto.OrganizerView{}, err
	}

	return organizer, nil
}

func (orgSvc *service) UpdateOrganizer(req dto.UpdateOrganizerRequest) error {
	err := orgSvc.organizerRepo.UpdateOrganizer(req)
	if err != nil {
		return err
	}

	return nil
}
