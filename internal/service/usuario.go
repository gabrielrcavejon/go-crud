package service

import (
	"errors"
	"go-crud/internal/model"
	"go-crud/internal/repository"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
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
func (s *UsuarioService) CriarUsuario(u model.Usuario) (uint, error) {
	u.Email = strings.TrimSpace(u.Email)
	u.Nome = strings.TrimSpace(u.Nome)

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return 0, errors.New("e-mail inválido")
	}

	if u.Nome == "" {
		return 0, errors.New("nome obrigatório")
	}

	if u.Senha == "" {
		return 0, errors.New("senha obrigatória")
	}

	hash, erro := bcrypt.GenerateFromPassword([]byte(u.Senha), bcrypt.DefaultCost)
	if erro != nil {
		return 0, errors.New("erro ao gerar hash da senha")
	}
	u.Senha = string(hash)

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
