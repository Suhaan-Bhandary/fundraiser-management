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
		orgDetail.Name, orgDetail.Detail, orgDetail.Email, orgDetail.Password, orgDetail.Mobile,
	)

	if err != nil {
		// Check if the error is a duplicate entry error
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			return internal_errors.DuplicateKeyError{
				Message: "Organizer with email already exists",
			}
		}

		return errors.New("error while creating the organizer")
	}

	return nil
}

func (organizerStore *organizerStore) GetOrganizerIDPasswordAndVerifyStatus(email string) (uint, string, bool, error) {
	var id uint
	var password string
	var isVerified bool

	row := organizerStore.db.QueryRow(getOrganizerIdPasswordQuery, email)
	err := row.Scan(&id, &password, &isVerified)

	if err != nil {
		return 0, "", false, internal_errors.NotFoundError{Message: "Invalid email or password"}
	}

	return id, password, isVerified, nil
}

func (organizerStore *organizerStore) VerifyOrganizer(organizerId uint) error {
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
		rows.Scan(&organizer.ID, &organizer.Name, &organizer.Detail, &organizer.Email, &organizer.Mobile, &organizer.IsVerified)
		organizers = append(organizers, organizer)
	}
	defer rows.Close()

	return organizers, nil
}

func (organizerStore *organizerStore) DeleteOrganizer(organizerId uint) error {
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

func (organizerStore *organizerStore) GetOrganizer(organizerId uint) (dto.OrganizerView, error) {
	var organizer dto.OrganizerView
	err := organizerStore.db.QueryRow(getOrganizerQuery, organizerId).Scan(
		&organizer.ID, &organizer.Name,
		&organizer.Detail, &organizer.Email,
		&organizer.Mobile, &organizer.IsVerified,
	)

	if err != nil {
		fmt.Println(err)
		return dto.OrganizerView{}, internal_errors.NotFoundError{Message: "Organizer not found"}
	}

	return organizer, nil
}

func (organizerStore *organizerStore) UpdateOrganizer(req dto.UpdateOrganizerRequest) error {
	res, err := organizerStore.db.Exec(
		updateOrganizerQuery,
		req.Email, req.Detail, req.Mobile, req.OrganizerId,
	)

	if err != nil {
		fmt.Println(err)
		// Check if the error is a duplicate entry error
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			return internal_errors.DuplicateKeyError{
				Message: "Update failed, email already in use",
			}
		}
		return errors.New("error while updating organizer")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return errors.New("error while updating organizer")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "Organizer not found"}
	}

	return nil
}
