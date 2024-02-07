package repository

import "time"

type FundraiserStorer interface{}

type Fundraiser struct {
	ID           uint
	Title        string
	Description  string
	OrganizerId  uint
	ImageUrl     string
	VideoUrl     string
	TargetAmount float64
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
