package repository

import (
	"database/sql"
	"go-crud/internal/model"
)

// UsuarioRepository é o modelo de usuario para o repository
type UsuarioRepository struct {
	DB *sql.DB
}

// NewUsuarioRepository instancia o novo repository a ser usado
func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

// Create vai criar um novo usuario no banco de dados
func (r *UsuarioRepository) Create(u model.Usuario) error {
	stmt, erro := r.DB.Prepare("INSERT INTO usuario(nome, email, senha) VALUES(?, ?, ?)")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(u.Nome, u.Email, u.Senha)
	return erro
}
