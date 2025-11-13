package handlers

import (
	"encoding/json"
	"go-crud/internal/dto"
	"go-crud/internal/response"
	"go-crud/internal/service"
	"net/http"
)

// LoginHandler e o handle, a estrutura
type LoginHandler struct {
	Service *service.LoginService
}

// NewLoginHandler retorna um novo handler
func NewLoginHandler(s *service.LoginService) *LoginHandler {
	return &LoginHandler{Service: s}
}

// Login e a rota onde vai poder ser eito autenticacao
func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest

	if erro := json.NewDecoder(r.Body).Decode(&request); erro != nil {
		response.RetonarErro(w, http.StatusBadRequest, "Falha ao ler Json")
		return
	}

	if erro := h.Service.Login(request.Email, request.Senha); erro != nil {
		response.RetonarErro(w, http.StatusUnauthorized, erro.Error())
		return
	}

	response.RetonarSucesso(w, http.StatusOK, nil, "")
}
