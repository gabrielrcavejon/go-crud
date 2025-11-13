package routes

import (
	"go-crud/internal/handlers"
	"go-crud/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterProdutoRoutes retorna as rotas de produto configuradas
func RegisterProdutoRoutes(router *mux.Router, handler *handlers.ProdutoHandler) {
	r := router.PathPrefix("/produto").Subrouter()
	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("", handler.CriarProduto).Methods(http.MethodPost)
	r.HandleFunc("", handler.GetProdutos).Methods(http.MethodGet)
	r.HandleFunc("/{idProduto}", handler.GetProduto).Methods(http.MethodGet)
	r.HandleFunc("/{idProduto}", handler.AtualizarProduto).Methods(http.MethodPut)
	r.HandleFunc("/{idProduto}", handler.DeletarProduto).Methods(http.MethodDelete)
}
