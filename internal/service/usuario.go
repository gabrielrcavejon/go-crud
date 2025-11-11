package service

import (
	"go-crud/internal/model"
	"go-crud/internal/repository"
)

// UsuarioService é o servico de usuario
type UsuarioService struct {
	Repo *repository.UsuarioRepository
}

// NewUsuarioService faz um novo servico de usuario
func NewUsuarioService(r *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{Repo: r}
}

// CriarUsuario cria um usuario no banco de dados
func (s *UsuarioService) CriarUsuario(u model.Usuario) error {
	// fazer senha com hasha qui
	return s.Repo.Create(u)
}

// DeleteUsuario deleta o usuario com id passado por parametro
func (s *UsuarioService) DeleteUsuario(idUsuario uint) error {
	return s.Repo.Delete(idUsuario)
}

// AtualizarUsuario atualiza o usuario
func (s *UsuarioService) AtualizarUsuario(idUsuario uint, u model.Usuario) error {
	return s.Repo.Update(idUsuario, u)
}

// GetUsuario pega um usuario unico pelo id
func (s *UsuarioService) GetUsuario(idUsuario uint) (model.Usuario, error) {
	return s.Repo.GetUsuario(idUsuario)
}

// GetUsuarios pega todos os usuarios do banco
func (s *UsuarioService) GetUsuarios() ([]model.Usuario, error) {
	return s.Repo.GetUsuarios()
}
