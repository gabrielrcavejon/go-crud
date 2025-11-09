package repository

import (
	"database/sql"
	"go-crud/internal/model"
)

// ProdutoRepository é o modelo de produto para o repository
type ProdutoRepository struct {
	DB *sql.DB
}

// NewProdutoRepository instancia o novo repository a ser usado
func NewProdutoRepository(db *sql.DB) *ProdutoRepository {
	return &ProdutoRepository{DB: db}
}

// Create vai criar um novo produto no banco de dados
func (r *ProdutoRepository) Create(p model.Produto) error {
	stmt, erro := r.DB.Prepare("INSERT INTO produto(nome, descricao, idusuario) VALUES(?, ?, ?)")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(p.Nome, p.Descricao, 1) // Por enquanto fixo
	return erro
}
