package handlers

import (
	"encoding/json"
	"go-crud/internal/model"
	"go-crud/internal/service"
	"net/http"
)

// ProdutoHandler é o handler de produto
type ProdutoHandler struct {
	Service *service.ProdutoService
}

// NewProdutoHandler cria um novo handler de produto
func NewProdutoHandler(s *service.ProdutoService) *ProdutoHandler {
	return &ProdutoHandler{Service: s}
}

// CriarProduto faz a criação de um novo produto
func (h *ProdutoHandler) CriarProduto(w http.ResponseWriter, r *http.Request) {
	var p model.Produto
	if erro := json.NewDecoder(r.Body).Decode(&p); erro != nil {
		http.Error(w, "JSON inválido: "+erro.Error(), http.StatusBadRequest)
		return
	}

	if erro := h.Service.CriarProduto(p); erro != nil {
		http.Error(w, "Erro ao criar produto: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Produto criado"})
}
