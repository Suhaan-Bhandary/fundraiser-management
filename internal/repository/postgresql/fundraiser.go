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

func (fundStore *fundraiserStore) GetFundraiser(fundraiserId int) (dto.FundraiserView, error) {
	var fundraiser dto.FundraiserView
	row := fundStore.db.QueryRow(getFundraiserQuery, fundraiserId)

	// select title, description, organizer_id, image_url, video_url, target_amount, status, 1 as organizer_name
	err := row.Scan(
		&fundraiser.ID, &fundraiser.Title, &fundraiser.Description, &fundraiser.OrganizerId,
		&fundraiser.OrganizerName, &fundraiser.ImageUrl, &fundraiser.VideoUrl, &fundraiser.TargetAmount,
		&fundraiser.Status, &fundraiser.CreatedAt, &fundraiser.UpdatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return dto.FundraiserView{}, internal_errors.NotFoundError{Message: "Fundraiser not found"}
	}

	return fundraiser, nil
}

func (fundStore *fundraiserStore) CloseFundraiser(fundraiserId int) error {
	res, err := fundStore.db.Exec(
		closeFundraiserQuery,
		fundraiserId,
	)

	if err != nil {
		return errors.New("error while closing the fundraiser")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error while closing the fundraiser")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "Fundraiser not found"}
	}

	return nil
}

func (fundStore *fundraiserStore) GetFundraiserOrganizerIdAndStatus(fundraiserId int) (int, string, error) {
	var organizerId int
	var fundraiserStatus string
	row := fundStore.db.QueryRow(getOrganizerIdAndStatusFromFundraiserQuery, fundraiserId)
	err := row.Scan(&organizerId, &fundraiserStatus)

	if err != nil {
		return -1, "", internal_errors.NotFoundError{Message: "Invalid username or password"}
	}

	return organizerId, fundraiserStatus, nil
}

func (fundStore *fundraiserStore) BanFundraiser(fundraiserId int) error {
	res, err := fundStore.db.Exec(
		banFundraiserQuery,
		fundraiserId,
	)

	if err != nil {
		return errors.New("error while banning the fundraiser")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error while banning the fundraiser")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "Fundraiser not found"}
	}

	return nil
}

func (fundraiserStore *fundraiserStore) ListFundraiser() ([]dto.FundraiserView, error) {
	fundraiserDetailList := []dto.FundraiserView{}

	rows, err := fundraiserStore.db.Query(listFundraisers)
	if err != nil {
		fmt.Println(err)
		return []dto.FundraiserView{}, errors.New("error while fetching user fundraiser")
	}
	defer rows.Close()

	for rows.Next() {
		var fundraiserDetail dto.FundraiserView
		err := rows.Scan(
			&fundraiserDetail.ID, &fundraiserDetail.Title, &fundraiserDetail.Description, &fundraiserDetail.OrganizerId,
			&fundraiserDetail.OrganizerName, &fundraiserDetail.ImageUrl, &fundraiserDetail.VideoUrl, &fundraiserDetail.TargetAmount,
			&fundraiserDetail.Status, &fundraiserDetail.CreatedAt, &fundraiserDetail.UpdatedAt,
		)

		if err != nil {
			return []dto.FundraiserView{}, errors.New("error while fetching user fundraiser")
		}

		fundraiserDetailList = append(fundraiserDetailList, fundraiserDetail)
	}

	return fundraiserDetailList, nil
}

func (fundStore *fundraiserStore) UpdateFundraiser(updateDetail dto.UpdateFundraiserRequest) error {
	res, err := fundStore.db.Exec(
		updateFundraiserQuery,
		updateDetail.Title, updateDetail.Description, updateDetail.ImageUrl, updateDetail.VideoUrl, updateDetail.TargetAmount, updateDetail.FundraiserId,
	)

	if err != nil {
		return errors.New("error while updating the fundraiser")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error while updating the fundraiser")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "Fundraiser not found"}
	}

	return nil
}
