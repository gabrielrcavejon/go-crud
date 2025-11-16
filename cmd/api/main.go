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

	"github.com/joho/godotenv"
)

func main() {
	// Carrega o .ENV
	erro := godotenv.Load("../../.env")
	if erro != nil {
		log.Fatal(".env não encontrado")
	}

	// Conecta no banco
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

	// Configura as rotas da api
	router := routes.Setup(usuarioHandler, produtoHandler, loginHandler)

	// EM AMBIENTE DE DESENVOLVIMENTO CASO FOR TESTAR COM DISPOSITIVOS DIFERENTES
	// DESCOMENTE ESSA PARTE DO CORS, MAS LEMBRE QUE ELE PERMITE QUALQUER
	// DISPOSITIVO ACESSAR SE HABILITAR ISSO

	/*c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true, // Descomente para log de debug
	})

	handler := c.Handler(router)*/

	// Start na api
	fmt.Println("API em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
