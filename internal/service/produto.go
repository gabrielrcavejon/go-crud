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
func (s *ProdutoService) CriarProduto(p model.Produto) (uint, error) {
	return s.Repo.Create(p)
}

// DeleteProduto deleta o produto com id passado por parametro
func (s *ProdutoService) DeleteProduto(idProduto uint) error {
	return s.Repo.Delete(idProduto)
}

// AtualizarProduto atualiza o produto
func (s *ProdutoService) AtualizarProduto(idProduto uint, p model.Produto) error {
	return s.Repo.Update(idProduto, p)
}

// GetProduto pega um produto unico pelo id
func (s *ProdutoService) GetProduto(idProduto uint) (model.Produto, error) {
	return s.Repo.GetProduto(idProduto)
}

// GetProdutos pega todos os produtos do banco
func (s *ProdutoService) GetProdutos() ([]model.Produto, error) {
	return s.Repo.GetProdutos()
}
