package dto

type Token struct {
	ID   int
	Role string
}

type MessageResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
