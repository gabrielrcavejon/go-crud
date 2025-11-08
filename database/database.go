package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conectar faz a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	conexao := "root:1234@/go?charset=utf8&parseTime=True&loc=Local"

	database, erro := sql.Open("mysql", conexao)

	if erro != nil {
		return nil, erro
	}

	// Tenta dar um ping
	if erro = database.Ping(); erro != nil {
		return nil, erro
	}

	return database, nil
}
