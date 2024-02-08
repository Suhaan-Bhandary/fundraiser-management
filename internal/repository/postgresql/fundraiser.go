package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

type fundraiserStore struct {
	db *sql.DB
}

func NewFundraiserRepo(db *sql.DB) repository.FundraiserStorer {
	return &fundraiserStore{
		db: db,
	}
}

func (fundStore *fundraiserStore) CreateFundraiser(fundDetail dto.CreateFundraiserRequest) error {
	_, err := fundStore.db.Exec(
		insertFundraiserQuery,
		fundDetail.Title, fundDetail.Description, fundDetail.OrganizerId, fundDetail.ImageUrl, fundDetail.VideoUrl, fundDetail.TargetAmount, constants.ACTIVE_STATUS,
	)

	if err != nil {
		fmt.Println(err)
		return errors.New("error while creating the fund")
	}

	return nil
}
