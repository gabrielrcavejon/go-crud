package main

import (
	"fmt"
	produto "go-crud/service"
	"log"
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

func main() {
	router := mux.NewRouter()
	router.Use(logMiddleware)

	produtoRouter := router.PathPrefix("/produto").Subrouter()
	{
		produtoRouter.HandleFunc("", produto.CriarProduto).Methods(http.MethodPost)
		produtoRouter.HandleFunc("", produto.GetProdutos).Methods(http.MethodGet)
		produtoRouter.HandleFunc("/{idProduto}", produto.GetProduto).Methods(http.MethodGet)
		produtoRouter.HandleFunc("/{idProduto}", produto.AtualizarProduto).Methods(http.MethodPut)
		produtoRouter.HandleFunc("/{idProduto}", produto.DeletarProduto).Methods(http.MethodDelete)
	}

	fmt.Println("API em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
