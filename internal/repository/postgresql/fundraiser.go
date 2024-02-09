package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
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

func (fundStore *fundraiserStore) CreateFundraiser(fundDetail dto.CreateFundraiserRequest) (int, error) {
	var fundraiserId int
	err := fundStore.db.QueryRow(
		insertFundraiserQuery,
		fundDetail.Title, fundDetail.Description, fundDetail.OrganizerId, fundDetail.ImageUrl, fundDetail.VideoUrl, fundDetail.TargetAmount, constants.ACTIVE_STATUS,
	).Scan(&fundraiserId)

	if err != nil {
		fmt.Println(err)
		return -1, errors.New("error while creating the fund")
	}

	return fundraiserId, nil
}

func (fundStore *fundraiserStore) DeleteFundraiser(fundraiserId int) error {
	res, err := fundStore.db.Exec(
		deleteFundraiserQuery,
		fundraiserId,
	)

	if err != nil {
		return errors.New("error while deleting the fundraiser")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error while deleting the fundraiser")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "Fundraiser not found"}
	}

	return nil
}

func (fundStore *fundraiserStore) GetFundraiserOrganizerId(fundraiserId int) (int, error) {
	var organizerId int
	row := fundStore.db.QueryRow(getOrganizerIdFromFundraiser, fundraiserId)
	err := row.Scan(&organizerId)

	if err != nil {
		return -1, internal_errors.NotFoundError{Message: "Invalid username or password"}
	}

	return organizerId, nil

}
