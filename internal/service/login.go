package service

import (
	"errors"
	"go-crud/internal/repository"
	"go-crud/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// LoginService e a estrutura do service do login
type LoginService struct {
	Repo *repository.UsuarioRepository
}

// NewLoginService faz um novo servico de login
func NewLoginService(r *repository.UsuarioRepository) *LoginService {
	return &LoginService{Repo: r}
}

// Login verifica se o usuario esta autorizado a usar a aplicacao
func (s *LoginService) Login(email string, senha string) (string, error) {
	retornoSenhaIncorreta := "Usuário ou senha incorretos"

	usuario, erro := s.Repo.GetUsuarioByEmail(email)
	if erro != nil || usuario.ID == 0 {
		return "", errors.New(retornoSenhaIncorreta)
	}

	if erro = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha)); erro != nil {
		return "", errors.New(retornoSenhaIncorreta)
	}

	token, erro := utils.GerarToken(uint(usuario.ID))
	if erro != nil {
		return "", errors.New(retornoSenhaIncorreta)
	}

	return token, nil
}
