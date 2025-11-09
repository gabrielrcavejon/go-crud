package service

import (
	"go-crud/internal/model"
	"go-crud/internal/repository"
)

// ProdutoService é o serviço de produto
type ProdutoService struct {
	Repo *repository.ProdutoRepository
}

// NewProdutoService cria um novo serviço de produto
func NewProdutoService(r *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{Repo: r}
}

// CriarProduto cria um produto no banco de dados
func (s *ProdutoService) CriarProduto(p model.Produto) error {
	return s.Repo.Create(p)
}
