package handlers

import (
	"encoding/json"
	"go-crud/internal/model"
	"go-crud/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UsuarioHandler e o handler de usuario
type UsuarioHandler struct {
	Service *service.UsuarioService
}

// NewUsuarioHandler faz um novo handler de usuario
func NewUsuarioHandler(s *service.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{Service: s}
}

// CriarUsuario faz a criacao de um novo usuario
func (h *UsuarioHandler) CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var u model.Usuario
	if erro := json.NewDecoder(r.Body).Decode(&u); erro != nil {
		http.Error(w, "JSON inválido: "+erro.Error(), http.StatusBadRequest)
		return
	}

	if erro := h.Service.CriarUsuario(u); erro != nil {
		http.Error(w, "Erro ao criar usuário: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensagem": "Usuário criado"})
}

// AtualizarUsario atualiza um usuario
func (h *UsuarioHandler) AtualizarUsario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idUsuario"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		http.Error(w, "ID do usuario inválido"+erro.Error(), http.StatusBadRequest)
		return
	}

	var u model.Usuario
	if erro := json.NewDecoder(r.Body).Decode(&u); erro != nil {
		http.Error(w, "JSON inválido: "+erro.Error(), http.StatusBadRequest)
		return
	}

	if erro := h.Service.AtualizarUsuario(uint(ID), u); erro != nil {
		http.Error(w, "Erro ao atualizar usuário: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario deleta um usuario
func (h *UsuarioHandler) DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idUsuario"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		http.Error(w, "ID do usuario inválido"+erro.Error(), http.StatusBadRequest)
		return
	}

	if erro := h.Service.DeleteUsuario(uint(ID)); erro != nil {
		http.Error(w, "Erro ao deletar usuário: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetUsuario pega o usuario a partir do id dele
func (h *UsuarioHandler) GetUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idUsuario"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		http.Error(w, "ID do usuario inválido"+erro.Error(), http.StatusBadRequest)
		return
	}

	usuario, erro := h.Service.GetUsuario(uint(ID))
	if erro != nil {
		http.Error(w, "Erro ao pegar um usuário: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]model.Usuario{"mensagem": usuario})
}

// GetUsuarios pega todos os usuarios do banco de dados
func (h *UsuarioHandler) GetUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, erro := h.Service.GetUsuarios()
	if erro != nil {
		http.Error(w, "Erro ao pegar todos os usuário: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]model.Usuario{"mensagem": usuarios})
}
