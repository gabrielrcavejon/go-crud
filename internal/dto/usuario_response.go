package dto

// UsuarioResponse e a estrutura que vai ser retornada no json
type UsuarioResponse struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
