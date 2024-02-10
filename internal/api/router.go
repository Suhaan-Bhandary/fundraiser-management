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
	userRouter.HandleFunc("", ListUsersHandler(deps.UserService)).Methods(http.MethodGet)
	userRouter.HandleFunc(
		"/{id}",
		middleware.CheckAuth(DeleteUserHandler(deps.UserService), []string{constants.ADMIN}),
	).Methods(http.MethodDelete)
	userRouter.HandleFunc(
		"/profile",
		middleware.CheckAuth(GetUserProfileHandler(deps.UserService), []string{constants.USER}),
	).Methods(http.MethodGet)
	userRouter.HandleFunc(
		"/donation",
		middleware.CheckAuth(ListUserDonationsHandler(deps.DonationService), []string{constants.USER}),
	).Methods(http.MethodGet)

	// Admin routes
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/login", LoginAdminHandler(deps.AdminService)).Methods(http.MethodPost)

	// Organizer routes
	organizerRouter := router.PathPrefix("/organizer").Subrouter()
	organizerRouter.HandleFunc("/register", RegisterOrganizerHandler(deps.OrganizerService)).Methods(http.MethodPost)
	organizerRouter.HandleFunc("/login", LoginOrganizerHandler(deps.OrganizerService)).Methods(http.MethodPost)
	organizerRouter.HandleFunc(
		"",
		middleware.CheckAuth(ListOrganizersHandler(deps.OrganizerService), []string{constants.ADMIN}),
	).Methods(http.MethodGet)
	organizerRouter.HandleFunc(
		"",
		middleware.CheckAuth(UpdateOrganizerHandler(deps.OrganizerService), []string{constants.ORGANIZER}),
	).Methods(http.MethodPut)
	organizerRouter.HandleFunc(
		"/{id}",
		GetOrganizerHandler(deps.OrganizerService),
	).Methods(http.MethodGet)
	organizerRouter.HandleFunc(
		"/{id}",
		middleware.CheckAuth(DeleteOrganizerHandler(deps.OrganizerService), []string{constants.ADMIN}),
	).Methods(http.MethodDelete)
	organizerRouter.HandleFunc(
		"/{id}/verify",
		middleware.CheckAuth(VerifyOrganizerHandler(deps.OrganizerService), []string{constants.ADMIN}),
	).Methods(http.MethodPatch)

	// Fundraiser routes
	fundraiserRouter := router.PathPrefix("/fundraiser").Subrouter()
	fundraiserRouter.HandleFunc("",
		middleware.CheckAuth(CreateFundraiserHandler(deps.FundraiserService), []string{constants.ORGANIZER}),
	).Methods(http.MethodPost)
	fundraiserRouter.HandleFunc("",
		ListFundraisersHandler(deps.FundraiserService),
	).Methods(http.MethodGet)
	fundraiserRouter.HandleFunc("/{id}", GetFundraiserHandler(deps.FundraiserService)).Methods(http.MethodGet)
	fundraiserRouter.HandleFunc("/{id}",
		middleware.CheckAuth(UpdateFundraiserHandler(deps.FundraiserService), []string{constants.ORGANIZER}),
	).Methods(http.MethodPut)
	fundraiserRouter.HandleFunc("/{id}",
		middleware.CheckAuth(DeleteFundraiserHandler(deps.FundraiserService), []string{constants.ORGANIZER, constants.ADMIN}),
	).Methods(http.MethodDelete)
	fundraiserRouter.HandleFunc("/{id}/donation",
		middleware.CheckAuth(CreateDonationHandler(deps.DonationService), []string{constants.USER}),
	).Methods(http.MethodPost)
	fundraiserRouter.HandleFunc("/{id}/donation",
		ListFundraiserDonationsHandler(deps.DonationService),
	).Methods(http.MethodGet)
	fundraiserRouter.HandleFunc("/{id}/close",
		middleware.CheckAuth(CloseFundraiserHandler(deps.FundraiserService), []string{constants.ORGANIZER}),
	).Methods(http.MethodPatch)
	fundraiserRouter.HandleFunc("/{id}/ban",
		middleware.CheckAuth(BanFundraiserHandler(deps.FundraiserService), []string{constants.ADMIN}),
	).Methods(http.MethodPatch)
	fundraiserRouter.HandleFunc("/{id}/unban",
		middleware.CheckAuth(UnBanFundraiserHandler(deps.FundraiserService), []string{constants.ADMIN}),
	).Methods(http.MethodPatch)

	// Donation list
	router.HandleFunc("/donation",
		middleware.CheckAuth(ListDonationsHandler(deps.DonationService), []string{constants.ADMIN}),
	).Methods(http.MethodGet)

	// Not Found Router
	router.HandleFunc("/", notFoundHandler).Methods(http.MethodGet)

	return router
}
