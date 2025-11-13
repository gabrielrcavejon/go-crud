package middleware

import (
	"fmt"
	"go-crud/internal/response"
	"go-crud/internal/utils"
	"net/http"
	"strings"
)

// AuthMiddleware verifica se a requisicao tem um token valido
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Pega o header Authorization
		auth := r.Header.Get("Authorization")
		if auth == "" {
			response.RetonarErro(w, http.StatusUnauthorized, "Token não informado")
			return
		}

		// Ve se o token e no padrao Bearer
		partes := strings.Split(auth, " ")
		if len(partes) != 2 || partes[0] != "Bearer" {
			response.RetonarErro(w, http.StatusUnauthorized, "Token inválido")
			return
		}

		tokenString := partes[1]

		// Valida o token
		claims, err := utils.ValidaToken(tokenString)
		if err != nil {
			response.RetonarErro(w, http.StatusUnauthorized, "Token expirado ou inválido")
			return
		}

		// Guarda o ID do usuario na request (caso precise)
		r.Header.Set("X-User-ID", fmt.Sprint(claims.IDUsuario))

		next.ServeHTTP(w, r)
	})
}
