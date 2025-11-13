package routes

import (
	"go-crud/internal/handlers"
	"go-crud/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterUsuarioRoutes retorna as rotas de usuario configuradas
func RegisterUsuarioRoutes(router *mux.Router, handler *handlers.UsuarioHandler) {
	r := router.PathPrefix("/usuario").Subrouter()
	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("", handler.CriarUsuario).Methods(http.MethodPost)
	r.HandleFunc("", handler.GetUsuarios).Methods(http.MethodGet)
	r.HandleFunc("/{idUsuario}", handler.GetUsuario).Methods(http.MethodGet)
	r.HandleFunc("/{idUsuario}", handler.AtualizarUsario).Methods(http.MethodPut)
	r.HandleFunc("/{idUsuario}", handler.DeletarUsuario).Methods(http.MethodDelete)
}
