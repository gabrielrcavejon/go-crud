package handlers

import (
	"encoding/json"
	"go-crud/internal/model"
	"go-crud/internal/service"
	"net/http"
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
