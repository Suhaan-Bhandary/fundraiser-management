package repository

import "time"

type AdminStorer interface {
	GetAdminIDPassword(username string) (uint, string, error)
}

type Admin struct {
	ID        uint
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
