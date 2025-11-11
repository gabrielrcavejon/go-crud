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

// Update vai atualizar um usuario
func (r *UsuarioRepository) Update(idUsuario uint, u model.Usuario) error {
	stmt, erro := r.DB.Prepare("UPDATE usuario SET nome = ?, email = ?, senha = ? WHERE idusuario = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(u.Nome, u.Email, u.Senha, idUsuario)
	return erro
}

// Delete vai deletar um usuario
func (r *UsuarioRepository) Delete(idUsuario uint) error {
	stmt, erro := r.DB.Prepare("DELETE FROM usuario WHERE idusuario = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(idUsuario)
	return erro
}

// GetUsuarios vai trazer todos os usaurios
func (r *UsuarioRepository) GetUsuarios() ([]model.Usuario, error) {
	rows, erro := r.DB.Query("SELECT idusuario, nome, email, senha FROM usuario")
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var usuarios []model.Usuario

	for rows.Next() {
		var u model.Usuario

		if erro := rows.Scan(&u.ID, &u.Nome, &u.Email, &u.Senha); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}

// GetUsuario vai trazer apenas o usuario com o id que foi passado como parametro
func (r *UsuarioRepository) GetUsuario(idUsuario uint) (model.Usuario, error) {
	row := r.DB.QueryRow("SELECT idusuario, nome, email, senha FROM usuario WHERE idusuario = ?", idUsuario)

	var u model.Usuario

	if erro := row.Scan(&u.ID, &u.Nome, &u.Email, &u.Senha); erro != nil {
		return model.Usuario{}, erro
	}

	return u, nil
}
