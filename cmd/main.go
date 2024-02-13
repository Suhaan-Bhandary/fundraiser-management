package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/api"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/admin"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	repository "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
	postgresql "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/postgresql"
	"github.com/joho/godotenv"
)

func main() {
	action := ""
	if len(os.Args) == 2 {
		action = os.Args[1]
	}

	// Setup env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	switch action {
	case "create-admin":
		createAdmin()
	default:
		startServer()
	}
}

func startServer() {
	ctx := context.Background()

	fmt.Println("Starting Server...")
	defer fmt.Println("Shutting Down Server...")

	// Initialize DB
	db, err := repository.InitializeDatabase(ctx)
	if err != nil {
		fmt.Println("DB Error:", err)
		return
	}

	// Creating services
	services := app.NewServices(db)

	// Initializing router
	router := api.NewRouter(services)

	// Listening to the server and assigning our custom router
	err = http.ListenAndServe(constants.SERVER_ADDRESS, router)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createAdmin() {
	ctx := context.Background()

	// Initialize DB
	db, err := repository.InitializeDatabase(ctx)
	adminRepo := postgresql.NewAdminRepo(db)
	adminService := admin.NewService(adminRepo)

	if err != nil {
		fmt.Println("DB Error:", err)
		return
	}

	var req dto.RegisterAdminRequest

	fmt.Print("Enter admin username: ")
	_, err = fmt.Scanln(&req.Username)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter admin password: ")
	_, err = fmt.Scanln(&req.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = adminService.RegisterAdmin(req)
	if err != nil {
		fmt.Println("Cannot create admin:", err)
		return
	}

	fmt.Println("Admin created successfully")
}
