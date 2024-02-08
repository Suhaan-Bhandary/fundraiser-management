package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type UserStorer interface {
	RegisterUser(userDetail dto.RegisterUserRequest) error
	GetUserIDPassword(email string) (int, string, error)
	DeleteUser(userId int) error
	GetUserList() ([]dto.UserView, error)
	GetUserProfile(userId int) (dto.UserView, error)
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
