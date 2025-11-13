package routes

import (
	"go-crud/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterLoginRoutes retorna as rotas de login configuradas
func RegisterLoginRoutes(router *mux.Router, handler *handlers.LoginHandler) {
	r := router.PathPrefix("/login").Subrouter()

	r.HandleFunc("", handler.Login).Methods(http.MethodPost)
}
