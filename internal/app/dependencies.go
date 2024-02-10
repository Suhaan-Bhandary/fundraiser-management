package app

import (
	"database/sql"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/admin"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/donation"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/fundraiser"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/organizer"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/user"
	repository "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/postgresql"
)

type Dependencies struct {
	UserService       user.Service
	AdminService      admin.Service
	OrganizerService  organizer.Service
	FundraiserService fundraiser.Service
	DonationService   donation.Service
}

func NewServices(db *sql.DB) Dependencies {
	userRepo := repository.NewUserRepo(db)
	adminRepo := repository.NewAdminRepo(db)
	organizerRepo := repository.NewOrganizerRepo(db)
	fundraiserRepo := repository.NewFundraiserRepo(db)
	donationRepo := repository.NewDonationRepo(db)

	userService := user.NewService(userRepo)
	adminService := admin.NewService(adminRepo)
	organizerService := organizer.NewService(organizerRepo)
	fundraiserService := fundraiser.NewService(fundraiserRepo)
	donationService := donation.NewService(donationRepo, fundraiserRepo)

	return Dependencies{
		UserService:       userService,
		AdminService:      adminService,
		OrganizerService:  organizerService,
		FundraiserService: fundraiserService,
		DonationService:   donationService,
	}
}
