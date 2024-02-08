package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/api"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
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

func createAdmin() {
	ctx := context.Background()

	// Initialize DB
	db, err := repository.InitializeDatabase(ctx)
	if err != nil {
		fmt.Println("DB Error:", err)
		return
	}

	var username, password string

	fmt.Print("Enter admin username: ")
	_, err = fmt.Scanln(&username)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter admin password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println(err)
		return
	}

	hashedPassword, err := helpers.HashPassword(password)
	db.Query("INSERT INTO admin (username, password) values($1, $2)", username, hashedPassword)
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
