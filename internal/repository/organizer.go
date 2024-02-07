package repository

import "time"

type OrganizerStorer interface{}

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
