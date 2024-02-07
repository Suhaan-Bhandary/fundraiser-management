package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/api"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	fmt.Println("Starting Server...")
	defer fmt.Println("Shutting Down Server...")

	// Setup env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

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
