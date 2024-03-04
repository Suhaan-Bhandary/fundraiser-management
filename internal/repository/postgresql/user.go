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

func (userStore *userStore) GetUserIDPassword(email string) (uint, string, error) {
	var id uint
	var password string

	row := userStore.db.QueryRow(getUserPasswordQuery, email)
	err := row.Scan(&id, &password)

	if err != nil {
		return 0, "", internal_errors.NotFoundError{Message: "Invalid username or password"}
	}

	return id, password, nil
}

func (userStore *userStore) DeleteUser(userId uint) error {
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

func (userStore *userStore) ListUsers(req dto.ListUserRequest) ([]dto.UserView, error) {
	toSkip := req.Offset * req.Limit
	rows, err := userStore.db.Query(
		listUsersQuery,
		req.Search, req.OrderByKey, req.OrderByIsAscending, toSkip, req.Limit,
	)

	if err != nil {
		fmt.Println(err)
		return []dto.UserView{}, errors.New("error while fetching users")
	}

	users := []dto.UserView{}
	for rows.Next() {
		user := dto.UserView{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return []dto.UserView{}, errors.New("error while fetching users")
		}

		users = append(users, user)
	}
	defer rows.Close()

	return users, nil
}

func (userStore *userStore) GetListUsersCount(req dto.ListUserRequest) (uint, error) {
	var count uint
	err := userStore.db.QueryRow(getListUsersCountQuery, req.Search).Scan(&count)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("error while fetching users")
	}

	return count, nil
}

func (userStore *userStore) GetUserProfile(userId uint) (dto.UserView, error) {
	row := userStore.db.QueryRow(getUserQuery, userId)

	user := dto.UserView{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return dto.UserView{}, internal_errors.InvalidCredentialError{Message: "Invalid user token"}
		}

		return dto.UserView{}, errors.New("error while fetching users")
	}

	return user, nil
}
