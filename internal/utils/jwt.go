package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims é a estrutura do JWT
type JWTClaims struct {
	IDUsuario uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GerarToken gera um JWT token a partir de um idusuario
func GerarToken(IDUsuario uint) (string, error) {
	claims := JWTClaims{
		IDUsuario: IDUsuario,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("CHAVE_JWT")))
}

// ValidaToken Serve para validar se o token esta certo
func ValidaToken(tokenString string) (*JWTClaims, error) {
	token, erro := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("CHAVE_JWT")), nil
	})
	if erro != nil {
		return nil, erro
	}

	// type assertion (assertiva de tipo)
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, erro
}

// PegaUserID o id do usuario direto do header ja lido no middleware
func PegaUserID(r *http.Request) string {
	return r.Header.Get("X-User-ID")
}
