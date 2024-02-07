package repository

import "time"

type AdminStorer interface{}

type Admin struct {
	ID        uint
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
