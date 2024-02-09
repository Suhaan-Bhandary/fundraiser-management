package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type OrganizerStorer interface {
	RegisterOrganizer(orgDetail dto.RegisterOrganizerRequest) error
	GetOrganizerIDPassword(email string) (int, string, error)
	VerifyOrganizer(organizerId int) error
	GetOrganizerList(search string, verified string) ([]dto.OrganizerView, error)
	DeleteOrganizer(organizerId int) error
	GetOrganizer(organizerId int) (dto.OrganizerView, error)
}

type Organizer struct {
	ID           uint
	Organization string
	Detail       string
	Email        string
	Password     string
	Mobile       string
	IsVerified   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
