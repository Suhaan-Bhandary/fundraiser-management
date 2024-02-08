package app

import (
	"database/sql"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/admin"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/user"
	repository "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/postgresql"
)

type Dependencies struct {
	UserService  user.Service
	AdminService admin.Service
}

func NewServices(db *sql.DB) Dependencies {
	userRepo := repository.NewUserRepo(db)
	adminRepo := repository.NewAdminRepo(db)

	userService := user.NewService(userRepo)
	adminService := admin.NewService(adminRepo)

	return Dependencies{
		UserService:  userService,
		AdminService: adminService,
	}
}
