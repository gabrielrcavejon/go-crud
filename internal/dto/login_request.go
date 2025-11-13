package dto

// LoginRequest e a estrutura que vai chegar na api
type LoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
