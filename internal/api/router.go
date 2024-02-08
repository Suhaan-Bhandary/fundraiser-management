package api

import (
	"errors"
	"net/http"

	"github.com/Suhaan-Bhandary/fundraiser-management/internal/app"
	"github.com/Suhaan-Bhandary/fundraiser-management/internal/pkg/constants"
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
	userRouter.HandleFunc(
		"/{id}",
		middleware.CheckAuth(DeleteUserHandler(deps.UserService), []string{constants.ADMIN}),
	).Methods(http.MethodDelete)

	// Admin routes
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/login", LoginAdminHandler(deps.AdminService)).Methods(http.MethodPost)

	// Organizer routes
	organizerRouter := router.PathPrefix("/organizer").Subrouter()
	organizerRouter.HandleFunc("/register", RegisterOrganizerHandler(deps.OrganizerService)).Methods(http.MethodPost)
	organizerRouter.HandleFunc("/login", LoginOranizerHandler(deps.OrganizerService)).Methods(http.MethodPost)

	// Fundraiser routes
	fundraiserRouter := router.PathPrefix("/fundraiser").Subrouter()
	fundraiserRouter.HandleFunc("", CreateFundraiserHandler(deps.FundraiserService)).Methods(http.MethodPost)

	// Not Found Router
	router.HandleFunc("/", notFoundHandler).Methods(http.MethodGet)

	return router
}
