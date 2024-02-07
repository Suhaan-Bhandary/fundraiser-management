package app

import (
	"database/sql"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/user"
	repository "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/postgresql"
)

type Dependencies struct {
	UserService user.Service
}

func NewServices(db *sql.DB) Dependencies {
	userRepo := repository.NewUserRepo(db)
	userService := user.NewService(userRepo)

	return Dependencies{
		UserService: userService,
	}
}
