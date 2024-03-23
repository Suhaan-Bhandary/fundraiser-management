package dto

type Token struct {
	ID   uint
	Role string
}

type MessageResponse struct {
	Message string `json:"message"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type OrganizerLoginResponse struct {
	Token       string `json:"token"`
	OrganizerId uint   `json:"organizer_id"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}
