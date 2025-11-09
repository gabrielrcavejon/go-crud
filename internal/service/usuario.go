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
