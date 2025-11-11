package handlers

import (
	"encoding/json"
	"go-crud/internal/model"
	"go-crud/internal/response"
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
		response.RetonarErro(w, http.StatusBadRequest, "JSON inválido: "+erro.Error())
		return
	}

	ID, erro := h.Service.CriarProduto(p)
	if erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao criar produto: "+erro.Error())
		return
	}

	p.ID = uint32(ID)

	response.RetonarSucesso(w, http.StatusCreated, p, "Produto criado")
}

// AtualizarProduto atualiza um produto
func (h *ProdutoHandler) AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idProduto"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "ID do produto inválido"+erro.Error())
		return
	}

	var p model.Produto
	if erro := json.NewDecoder(r.Body).Decode(&p); erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "JSON inválido: "+erro.Error())
		return
	}

	if erro := h.Service.AtualizarProduto(uint(ID), p); erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao atualizar produto: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusNoContent, nil, "Produto Atualizado")
}

// DeletarProduto deleta um produto
func (h *ProdutoHandler) DeletarProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idProduto"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "ID do produto inválido"+erro.Error())
		return
	}

	if erro := h.Service.DeleteProduto(uint(ID)); erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao deletar produto: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusNoContent, nil, "Produto Excluido")
}

// GetProduto pega o produto a partir do id dele
func (h *ProdutoHandler) GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idProduto"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "ID do produto inválido"+erro.Error())
		return
	}

	produto, erro := h.Service.GetProduto(uint(ID))
	if erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao pedar um produto: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusOK, produto, "Produto encontrado")
}

// GetProdutos pega todos os produtos do banco de dados
func (h *ProdutoHandler) GetProdutos(w http.ResponseWriter, r *http.Request) {
	produtos, erro := h.Service.GetProdutos()
	if erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao pegar todos os produtos: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusOK, produtos, "Produtos listados")
}
