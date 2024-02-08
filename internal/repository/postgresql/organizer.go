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

func (organizerStore *organizerStore) GetOrganizerList(search string, verified string) ([]dto.OrganizerView, error) {
	var rows *sql.Rows
	var err error

	// TODO: Check how to use verified also here
	if len(search) >= 3 {
		rows, err = organizerStore.db.Query(getOrganizersWithFilter, search)
	} else {
		rows, err = organizerStore.db.Query(getOrganizers)
	}

	if err != nil {
		fmt.Println(err)
		return []dto.OrganizerView{}, errors.New("error while fetching organizers")
	}

	organizers := []dto.OrganizerView{}
	for rows.Next() {
		organizer := dto.OrganizerView{}
		rows.Scan(&organizer.ID, &organizer.Organization, &organizer.Detail, &organizer.Email, &organizer.Mobile, &organizer.IsVerified)
		organizers = append(organizers, organizer)
	}
	defer rows.Close()

	return organizers, nil
}

func (organizerStore *organizerStore) DeleteOrganizer(organizerId int) error {
	res, err := organizerStore.db.Exec(
		deleteOrganizerQuery,
		organizerId,
	)

	if err != nil {
		return errors.New("error while deleting the organizer")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error while deleting the organizer")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "Organizer not found"}
	}

	return nil
}
