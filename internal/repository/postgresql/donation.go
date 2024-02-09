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
