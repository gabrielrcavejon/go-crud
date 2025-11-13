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

	// USUARIO
	usuarioRepo := repository.NewUsuarioRepository(database.DB)
	usuarioService := service.NewUsuarioService(usuarioRepo)
	usuarioHandler := handlers.NewUsuarioHandler(usuarioService)

	// PRODUTO
	produtoRepo := repository.NewProdutoRepository(database.DB)
	produtoService := service.NewProdutoService(produtoRepo)
	produtoHandler := handlers.NewProdutoHandler(produtoService)

	// LOGIN
	loginService := service.NewLoginService(usuarioRepo)
	loginHandler := handlers.NewLoginHandler(loginService)

	router := routes.Setup(usuarioHandler, produtoHandler, loginHandler)

	fmt.Println("API em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
