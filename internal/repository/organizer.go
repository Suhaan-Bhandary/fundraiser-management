package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type OrganizerStorer interface {
	RegisterOrganizer(orgDetail dto.RegisterOrganizerRequest) error
	GetOrganizerIDPasswordAndVerifyStatus(email string) (uint, string, bool, error)
	VerifyOrganizer(organizerId uint) error
	GetOrganizerList(search string, verified string) ([]dto.OrganizerView, error)
	DeleteOrganizer(organizerId uint) error
	GetOrganizer(organizerId uint) (dto.OrganizerView, error)
	UpdateOrganizer(req dto.UpdateOrganizerRequest) error
}

type Organizer struct {
	ID         uint
	Name       string
	Detail     string
	Email      string
	Password   string
	Mobile     string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
