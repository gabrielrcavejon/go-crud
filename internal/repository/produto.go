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

// Update atualiza um produto no banco de dados
func (r *ProdutoRepository) Update(idProduto uint32, p model.Produto) error {
	stmt, erro := r.DB.Prepare("UPDATE product SET nome = ?, descricao = ?, idusuario = ? WHERE idproduto = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(p.Nome, p.Descricao, 1, idProduto) // Por enquanto fixo o idUsuario
	return erro
}

// Delete vai deletar um produto no banco de dados
func (r *ProdutoRepository) Delete(idProduto uint32) error {
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
	rows, erro := r.DB.Query("SELECT idproduto, nome, descricao, idusuario")
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
	row := r.DB.QueryRow("SELECT idproduto, nome, descricao, idusuario WHERE idusuario = ?", idProduto)

	var p model.Produto

	if erro := row.Scan(&p.ID, &p.Nome, &p.Descricao, &p.IDUsuario); erro != nil {
		return model.Produto{}, erro
	}

	return p, nil
}
