package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Este é o Driver do MYSQL, que deve ser importado para ser usado indiretamente
	"github.com/joho/godotenv"
)

// DB é a variavel global que contem a conexao do bancod de dados, e foi instanciada apenas uma vez
var DB *sql.DB

// Conectar e a funcao que vai fazer a conexao com o banco de dados
func Conectar() {
	erro := godotenv.Load("../../.env")
	if erro != nil {
		log.Fatal(".env não encontrado")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORTA"),
		os.Getenv("DB_NOME"),
	)

	DB, erro = sql.Open("mysql", dsn)
	if erro != nil {
		log.Fatalf("Erro ao conectar: %v", erro)
	}

	erro = DB.Ping()
	if erro != nil {
		log.Fatalf("Erro no ping com banco: %v", erro)
	}

	log.Println("Sucesso ao conectar com banco!")
}
