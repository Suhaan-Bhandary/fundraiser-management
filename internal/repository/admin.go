package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type AdminStorer interface {
	RegisterAdmin(req dto.RegisterAdminRequest) error
	GetAdminIDPassword(username string) (uint, string, error)
}

type Admin struct {
	ID        uint
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
