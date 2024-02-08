package api

import (
	"errors"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/middleware"
	"github.com/gorilla/mux"
)

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	error := errors.New("API route not found")
	middleware.ErrorResponse(w, 404, error)
}

func NewRouter(deps app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	// user routes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/register", RegisterUserHandler(deps.UserService)).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", LoginUserHandler(deps.UserService)).Methods(http.MethodPost)
	userRouter.HandleFunc("", UserListHandler(deps.UserService)).Methods(http.MethodGet)

	// Admin routes
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/login", LoginAdminHandler(deps.AdminService)).Methods(http.MethodPost)

	// Not Found Router
	router.HandleFunc("/", notFoundHandler).Methods(http.MethodGet)

	return router
}
