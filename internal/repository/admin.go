package repository

import "time"

type AdminStorer interface {
	GetAdminIDPassword(username string) (int, string, error)
}

type Admin struct {
	ID        uint
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
