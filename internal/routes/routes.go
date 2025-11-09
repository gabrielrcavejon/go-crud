package routes

import (
	"fmt"
	"go-crud/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// Middleware para fazer um log
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Setup vai retornar todas as rotas feitas
func Setup(usuarioHandler *handlers.UsuarioHandler, produtoHandler *handlers.ProdutoHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(logMiddleware)

	produtoRouter := router.PathPrefix("/produto").Subrouter()

	produtoRouter.HandleFunc("", produtoHandler.CriarProduto).Methods(http.MethodPost)
	//produtoRouter.HandleFunc("", produtoHandler.GetProdutos).Methods(http.MethodGet)
	//produtoRouter.HandleFunc("/{idProduto}", produtoHandler.GetProduto).Methods(http.MethodGet)
	//produtoRouter.HandleFunc("/{idProduto}", produtoHandler.AtualizarProduto).Methods(http.MethodPut)
	//produtoRouter.HandleFunc("/{idProduto}", produtoHandler.DeletarProduto).Methods(http.MethodDelete)

	return router
}
