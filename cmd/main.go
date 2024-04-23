package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/api"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app/admin"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/dto"
	repository "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository"
	postgresql "github.com/Suhaan-Bhandary/fundraiser-management/internal/repository/postgresql"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	action := ""
	if len(os.Args) == 4 {
		action = os.Args[1]
	}

	// Setup env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	switch action {
	case "create-admin":
		username := os.Args[2]
		password := os.Args[3]
		createAdmin(username, password)
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

	// Cors
	allowedOrigin := os.Getenv("ORIGIN_ALLOWED")

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{allowedOrigin})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Listening to the server and assigning our custom router
	err = http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), handlers.CORS(credentials, methods, origins, headers)(router))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createAdmin(username, password string) {
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
	req.Username = username
	req.Password = password

	err = adminService.RegisterAdmin(req)
	if err != nil {
		fmt.Println("Cannot create admin:", err)
		return
	}

	fmt.Println("Admin created successfully")
}
