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

		return errors.New("error while creating the user")
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

func (userStore *userStore) DeleteUser(userId int) error {
	res, err := userStore.db.Exec(
		deleteUserQuery,
		userId,
	)
	if err != nil {
		return errors.New("error while deleting the user")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error while deleting the user")
	}

	if rowsAffected == 0 {
		return internal_errors.NotFoundError{Message: "User not found"}
	}

	return nil
}

func (userStore *userStore) GetUserList() ([]dto.UserView, error) {
	rows, err := userStore.db.Query(getUsersQuery)
	if err != nil {
		return []dto.UserView{}, errors.New("error while fetching users")
	}

	users := []dto.UserView{}
	for rows.Next() {
		user := dto.UserView{}
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		users = append(users, user)
	}
	defer rows.Close()

	return users, nil
}
