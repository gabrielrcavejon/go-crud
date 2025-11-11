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
		response.RetonarErro(w, http.StatusBadRequest, "JSON inválido: "+erro.Error())
		return
	}

	ID, erro := h.Service.CriarUsuario(u)

	if erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao criar usuário: "+erro.Error())
		return
	}

	u.ID = uint32(ID)

	response.RetonarSucesso(w, http.StatusCreated, u, "Usuário criado")
}

// AtualizarUsario atualiza um usuario
func (h *UsuarioHandler) AtualizarUsario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idUsuario"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "ID do usuario inválido"+erro.Error())
		return
	}

	var u model.Usuario
	if erro := json.NewDecoder(r.Body).Decode(&u); erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "JSON inválido: "+erro.Error())
		return
	}

	if erro := h.Service.AtualizarUsuario(uint(ID), u); erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao atualizar usuário: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusNoContent, nil, "Usuario atualizado")
}

// DeletarUsuario deleta um usuario
func (h *UsuarioHandler) DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idUsuario"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "ID do usuario inválido"+erro.Error())
		return
	}

	if erro := h.Service.DeleteUsuario(uint(ID)); erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao deletar usuário: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusNoContent, nil, "Usuario deletado")
}

// GetUsuario pega o usuario a partir do id dele
func (h *UsuarioHandler) GetUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idUsuario"]

	ID, erro := strconv.ParseUint(idStr, 10, 64)
	if erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "ID do usuario inválido"+erro.Error())
		return
	}

	usuario, erro := h.Service.GetUsuario(uint(ID))
	if erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao pegar um usuário: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusOK, usuario, "Usuario Encontrado")
}

// GetUsuarios pega todos os usuarios do banco de dados
func (h *UsuarioHandler) GetUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, erro := h.Service.GetUsuarios()
	if erro != nil {
		response.RetonarErro(w, http.StatusInternalServerError, "Erro ao pegar todos os usuário: "+erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusOK, usuarios, "Usuarios listado")
}
