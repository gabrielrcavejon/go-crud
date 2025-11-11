package main

import (
	"fmt"
	"go-crud/internal/database"
	"go-crud/internal/handlers"
	"go-crud/internal/repository"
	"go-crud/internal/routes"
	"go-crud/internal/service"
	"log"
	"net/http"
)

func main() {
	database.Conectar()

	// Faz a injeção de dependencia automatica
	usuarioRepo := repository.NewUsuarioRepository(database.DB)
	usuarioService := service.NewUsuarioService(usuarioRepo)
	usuarioHandler := handlers.NewUsuarioHandler(usuarioService)

	produtoRepo := repository.NewProdutoRepository(database.DB)
	produtoService := service.NewProdutoService(produtoRepo)
	produtoHandler := handlers.NewProdutoHandler(produtoService)

	router := routes.Setup(usuarioHandler, produtoHandler)

	fmt.Println("API em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
