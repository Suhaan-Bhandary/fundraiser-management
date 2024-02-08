package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
	"github.com/lib/pq"
)

type organizerStore struct {
	db *sql.DB
}

func NewOrganizerRepo(db *sql.DB) repository.OrganizerStorer {
	return &organizerStore{
		db: db,
	}
}

func (organizerStore *organizerStore) RegisterOrganizer(orgDetail dto.RegisterOrganizerRequest) error {
	_, err := organizerStore.db.Exec(
		insertOrganizerQuery,
		orgDetail.Organization, orgDetail.Detail, orgDetail.Email, orgDetail.Password, orgDetail.Mobile,
	)

	if err != nil {
		// Check if the error is a duplicate entry error
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			return internal_errors.DuplicateKeyError{
				Message: "Organization with email already exists",
			}
		}

		return errors.New("error while creating the organizer")
	}

	return nil
}

func (organizerStore *organizerStore) GetOrganizerIDPassword(email string) (int, string, error) {
	var id int
	var password string

	row := organizerStore.db.QueryRow(getOrganizerIdPasswordQuery, email)
	err := row.Scan(&id, &password)

	if err != nil {
		return -1, "", internal_errors.NotFoundError{Message: "Invalid email or password"}
	}

	return id, password, nil
}

func (organizerStore *organizerStore) VerifyOrganizer(organizerId int) error {
	res, err := organizerStore.db.Exec(
		verifyOrganizerQuery,
		organizerId,
	)

	if err != nil {
		fmt.Println(err)
		return errors.New("error while verifying the organizer")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return errors.New("error while verifying the organizer")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "Organizer not found"}
	}

	return nil
}
