package routes

import (
	"fmt"
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
func Setup() *mux.Router {
	router := mux.NewRouter()
	router.Use(logMiddleware)

	produto := router.PathPrefix("/produto").Subrouter()

	/*produto.HandleFunc("", produto.CriarProduto).Methods(http.MethodPost)
	produto.HandleFunc("", produto.GetProdutos).Methods(http.MethodGet)
	produto.HandleFunc("/{idProduto}", produto.GetProduto).Methods(http.MethodGet)
	produto.HandleFunc("/{idProduto}", produto.AtualizarProduto).Methods(http.MethodPut)
	produto.HandleFunc("/{idProduto}", produto.DeletarProduto).Methods(http.MethodDelete)*/

	return router
}
