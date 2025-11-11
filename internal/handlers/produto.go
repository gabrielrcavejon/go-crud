package handlers

import (
	"encoding/json"
	"go-crud/internal/model"
	"go-crud/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// AtualizarProduto atualiza um produto
func (h *ProdutoHandler) AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idProduto"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		http.Error(w, "ID do produto inválido"+erro.Error(), http.StatusBadRequest)
		return
	}

	var p model.Produto
	if erro := json.NewDecoder(r.Body).Decode(&p); erro != nil {
		http.Error(w, "JSON inválido: "+erro.Error(), http.StatusBadRequest)
		return
	}

	if erro := h.Service.AtualizarProduto(uint(ID), p); erro != nil {
		http.Error(w, "Erro ao atualizar produto: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarProduto deleta um produto
func (h *ProdutoHandler) DeletarProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idProduto"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		http.Error(w, "ID do produto inválido"+erro.Error(), http.StatusBadRequest)
		return
	}

	if erro := h.Service.DeleteProduto(uint(ID)); erro != nil {
		http.Error(w, "Erro ao deletar produto: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetProduto pega o produto a partir do id dele
func (h *ProdutoHandler) GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idProduto"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		http.Error(w, "ID do produto inválido"+erro.Error(), http.StatusBadRequest)
		return
	}

	produto, erro := h.Service.GetProduto(uint(ID))
	if erro != nil {
		http.Error(w, "Erro ao pedar um produto: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]model.Produto{"mensagem": produto})
}

// GetProdutos pega todos os produtos do banco de dados
func (h *ProdutoHandler) GetProdutos(w http.ResponseWriter, r *http.Request) {
	produtos, erro := h.Service.GetProdutos()
	if erro != nil {
		http.Error(w, "Erro ao pegar todos os produtos: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]model.Produto{"mensagem": produtos})
}
