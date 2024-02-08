package repository

import (
	"database/sql"
	"errors"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
	"github.com/lib/pq"
)

type userStore struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserStorer {
	return &userStore{
		db: db,
	}
}

func (userStore *userStore) RegisterUser(userDetail dto.RegisterUserRequest) error {
	_, err := userStore.db.Exec(
		insertUserQuery,
		userDetail.FirstName, userDetail.LastName, userDetail.Email, userDetail.Password,
	)

	if err != nil {
		// Check if the error is a duplicate entry error
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			return internal_errors.DuplicateKeyError{
				Message: "User with email already exists",
			}
		}

		return errors.New("Error while creating the user")
	}

	return nil
}

func (userStore *userStore) GetUserIDPassword(email string) (int, string, error) {
	var id int
	var password string

	row := userStore.db.QueryRow(getUserPasswordQuery, email)
	err := row.Scan(&id, &password)

	if err != nil {
		return -1, "", internal_errors.NotFoundError{Message: "Invalid username or password"}
	}

	return id, password, nil
}
