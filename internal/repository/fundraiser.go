package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type FundraiserStorer interface {
	CreateFundraiser(fundDetail dto.CreateFundraiserRequest) (int, error)
	DeleteFundraiser(fundraiserId int) error
	GetFundraiserOrganizerId(fundraiserId int) (int, error)
	GetFundraiser(fundraiserId int) (dto.FundraiserView, error)
	CloseFundraiser(fundraiserId int) error
	GetFundraiserOrganizerIdAndStatus(fundraiserId int) (int, string, error)
}

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
