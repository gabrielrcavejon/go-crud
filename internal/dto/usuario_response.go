package dto

// UsuarioResponse e a estrutura que vai ser retornada no json
type UsuarioResponse struct {
	IDUsuario uint32 `json:"idusuario"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
}
