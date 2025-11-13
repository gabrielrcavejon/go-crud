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
func Setup(usuarioHandler *handlers.UsuarioHandler, produtoHandler *handlers.ProdutoHandler, loginHandler *handlers.LoginHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(logMiddleware)

	RegisterProdutoRoutes(router, produtoHandler)

	RegisterUsuarioRoutes(router, usuarioHandler)

	RegisterLoginRoutes(router, loginHandler)

	return router
}
