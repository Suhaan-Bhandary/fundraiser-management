package repository

import (
	"database/sql"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/internal_errors"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
)

type adminStore struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) repository.AdminStorer {
	return &adminStore{
		db: db,
	}
}

func (adminStore *adminStore) GetAdminIDPassword(username string) (uint, string, error) {
	var id uint
	var password string

	row := adminStore.db.QueryRow(getAdminIdPasswordQuery, username)
	err := row.Scan(&id, &password)

	if err != nil {
		return 0, "", internal_errors.NotFoundError{Message: "Invalid username or password"}
	}

	return id, password, nil
}
