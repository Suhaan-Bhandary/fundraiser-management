package repository

import (
	"time"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
)

type FundraiserStorer interface {
	CreateFundraiser(fundDetail dto.CreateFundraiserRequest) (uint, error)
	DeleteFundraiser(fundraiserId uint) error
	GetFundraiserOrganizerId(fundraiserId uint) (uint, error)
	GetFundraiser(fundraiserId uint) (dto.FundraiserView, error)
	CloseFundraiser(fundraiserId uint) error
	BanFundraiser(fundraiserId uint) error
	UnBanFundraiser(fundraiserId uint) error
	GetFundraiserOrganizerIdAndStatus(fundraiserId uint) (uint, string, error)
	ListFundraiser() ([]dto.FundraiserView, error)
	UpdateFundraiser(updateDetail dto.UpdateFundraiserRequest) error
	GetFundraiserStatus(fundraiserId uint) (string, error)
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
