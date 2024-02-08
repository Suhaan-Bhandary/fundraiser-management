package app

import (
	"database/sql"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/admin"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/organizer"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/user"
	repository "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/postgresql"
)

type Dependencies struct {
	UserService      user.Service
	AdminService     admin.Service
	OrganizerService organizer.Service
}

func NewServices(db *sql.DB) Dependencies {
	userRepo := repository.NewUserRepo(db)
	adminRepo := repository.NewAdminRepo(db)
	organizerRepo := repository.NewOrganizerRepo(db)

	userService := user.NewService(userRepo)
	adminService := admin.NewService(adminRepo)
	organizerService := organizer.NewService(organizerRepo)

	return Dependencies{
		UserService:      userService,
		AdminService:     adminService,
		OrganizerService: organizerService,
	}
}
