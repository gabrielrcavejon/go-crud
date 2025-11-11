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
func (r *ProdutoRepository) Create(p model.Produto) (uint, error) {
	stmt, erro := r.DB.Prepare("INSERT INTO produto(nome, descricao, idusuario) VALUES(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer stmt.Close()

	res, erro := stmt.Exec(p.Nome, p.Descricao, p.IDUsuario)
	if erro != nil {
		return 0, erro
	}

	ID, erro := res.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint(ID), nil
}

// Update atualiza um produto no banco de dados
func (r *ProdutoRepository) Update(idProduto uint, p model.Produto) error {
	stmt, erro := r.DB.Prepare("UPDATE produto SET nome = ?, descricao = ?, idusuario = ? WHERE idproduto = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(p.Nome, p.Descricao, p.IDUsuario, idProduto)
	return erro
}

// Delete vai deletar um produto no banco de dados
func (r *ProdutoRepository) Delete(idProduto uint) error {
	stmt, erro := r.DB.Prepare("DELETE FROM produto WHERE idproduto = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(idProduto)
	return erro
}

// GetProdutos vai retornar uma lista de produtos
func (r *ProdutoRepository) GetProdutos() ([]model.Produto, error) {
	rows, erro := r.DB.Query("SELECT idproduto, nome, descricao, idusuario FROM produto")
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var produtos []model.Produto

	for rows.Next() {
		var p model.Produto

		if erro := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.IDUsuario); erro != nil {
			return nil, erro
		}

		produtos = append(produtos, p)
	}

	return produtos, nil
}

// GetProduto vai retornar uma produto em especifico, dependendo do id que for passado por parametro
func (r *ProdutoRepository) GetProduto(idProduto uint) (model.Produto, error) {
	row := r.DB.QueryRow("SELECT idproduto, nome, descricao, idusuario FROM produto WHERE idproduto = ?", idProduto)

	var p model.Produto

	if erro := row.Scan(&p.ID, &p.Nome, &p.Descricao, &p.IDUsuario); erro != nil {
		return model.Produto{}, erro
	}

	return p, nil
}
