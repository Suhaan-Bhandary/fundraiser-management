package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type UserStorer interface {
	RegisterUser(userDetail dto.RegisterUserRequest) error
	GetUserIDPassword(email string) (uint, string, error)
	DeleteUser(userId uint) error
	ListUsers() ([]dto.UserView, error)
	GetUserProfile(userId uint) (dto.UserView, error)
}

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
