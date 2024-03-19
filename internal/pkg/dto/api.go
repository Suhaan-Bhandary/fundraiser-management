package dto

type Token struct {
	ID   uint
	Role string
}

type MessageResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	OrganizerId uint   `json:"organizer_id"`
	Token       string `json:"token"`
}
