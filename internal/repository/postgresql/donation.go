package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

type donationStore struct {
	db *sql.DB
}

func NewDonationRepo(db *sql.DB) repository.DonationStorer {
	return &donationStore{
		db: db,
	}
}

func (donationStore *donationStore) CreateDonation(donationDetail dto.CreateDonationRequest) (int, error) {
	var donationId int
	err := donationStore.db.QueryRow(
		insertDonationQuery,
		donationDetail.UserId, donationDetail.FundraiserId, donationDetail.Amount, donationDetail.IsAnonymous,
	).Scan(&donationId)

	if err != nil {
		fmt.Println(err)
		return -1, errors.New("error while creating donation")
	}

	return donationId, nil
}

func (donationStore *donationStore) ListUserDonations(user_id int) ([]dto.DonationView, error) {
	donationDetailList := []dto.DonationView{}

	rows, err := donationStore.db.Query(listUserDonations, user_id)
	if err != nil {
		fmt.Println(err)
		return []dto.DonationView{}, errors.New("error while fetching user donation")
	}
	defer rows.Close()

	for rows.Next() {
		var donationDetail dto.DonationView
		err := rows.Scan(
			&donationDetail.ID, &donationDetail.FundraiserId, &donationDetail.FundraiserTitle, &donationDetail.Amount,
			&donationDetail.IsAnonymous, &donationDetail.CreatedAt,
		)

		if err != nil {
			return []dto.DonationView{}, errors.New("error while fetching user donation")
		}

		donationDetailList = append(donationDetailList, donationDetail)
	}

	return donationDetailList, nil
}

func (donationStore *donationStore) ListFundraiserDonations(fundraiserId int) ([]dto.FundraiserDonationView, error) {
	donationDetailList := []dto.FundraiserDonationView{}

	rows, err := donationStore.db.Query(listFundraiserDonationsQuery, fundraiserId)
	if err != nil {
		fmt.Println(err)
		return []dto.FundraiserDonationView{}, errors.New("error while fetching fundraiser donation")
	}
	defer rows.Close()

	for rows.Next() {
		var donationDetail dto.FundraiserDonationView
		err := rows.Scan(
			&donationDetail.ID, &donationDetail.UserId, &donationDetail.UserName,
			&donationDetail.FundraiserId,
			&donationDetail.FundraiserTitle, &donationDetail.Amount,
			&donationDetail.IsAnonymous, &donationDetail.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []dto.FundraiserDonationView{}, errors.New("error while fetching fundraiser donation")
		}

		donationDetailList = append(donationDetailList, donationDetail)
	}

	return donationDetailList, nil
}

func (donationStore *donationStore) ListDonations() ([]dto.FundraiserDonationView, error) {
	donationDetailList := []dto.FundraiserDonationView{}

	rows, err := donationStore.db.Query(listDonationsQuery)
	if err != nil {
		fmt.Println(err)
		return []dto.FundraiserDonationView{}, errors.New("error while fetching donations")
	}
	defer rows.Close()

	for rows.Next() {
		var donationDetail dto.FundraiserDonationView
		err := rows.Scan(
			&donationDetail.ID, &donationDetail.UserId, &donationDetail.UserName,
			&donationDetail.FundraiserId,
			&donationDetail.FundraiserTitle, &donationDetail.Amount,
			&donationDetail.IsAnonymous, &donationDetail.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []dto.FundraiserDonationView{}, errors.New("error while fetching donations")
		}

		donationDetailList = append(donationDetailList, donationDetail)
	}

	return donationDetailList, nil
}
