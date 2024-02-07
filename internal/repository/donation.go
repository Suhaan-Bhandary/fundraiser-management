package repository

import "time"

type DonationStorer interface{}

type Donation struct {
	ID           uint
	UserId       uint
	FundraiserId uint
	Amount       float64
	IsAnonymous  bool
	CreatedAt    time.Time
}
