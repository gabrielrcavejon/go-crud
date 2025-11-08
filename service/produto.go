package produto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-crud/database"

	"github.com/gorilla/mux"
)

type produto struct {
	IDproduto uint32 `json:"idproduto"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}

// ===== Helpers =====
func respostaJSON(w http.ResponseWriter, status int, dados any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if dados != nil {
		json.NewEncoder(w).Encode(dados)
	}
}

func respostaErro(w http.ResponseWriter, status int, msg string, erro error) {
	if erro != nil {
		msg = msg + ": " + erro.Error()
	}
	http.Error(w, msg, status)
}

// ===== Handlers =====

// CriarProduto cria um produto
func CriarProduto(w http.ResponseWriter, r *http.Request) {
	var p produto
	if erro := json.NewDecoder(r.Body).Decode(&p); erro != nil {
		respostaErro(w, http.StatusBadRequest, "JSON inválido", erro)
		return
	}

	banco, erro := database.Conectar()
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Falha ao conectar no banco", erro)
		return
	}
	defer banco.Close()

	stmt, erro := banco.Prepare("INSERT INTO produto(nome, descricao) VALUES(?, ?)")
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro no prepare", erro)
		return
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(p.Nome, p.Descricao); erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro ao inserir produto", erro)
		return
	}

	respostaJSON(w, http.StatusCreated, map[string]string{"mensagem": "Produto criado"})
}

// GetProdutos retorna todos os produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	banco, erro := database.Conectar()
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Falha ao conectar no banco", erro)
		return
	}
	defer banco.Close()

	rows, erro := banco.Query("SELECT idproduto, nome, descricao FROM produto")
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro ao buscar produtos", erro)
		return
	}
	defer rows.Close()

	var produtos []produto
	for rows.Next() {
		var p produto
		if erro := rows.Scan(&p.IDproduto, &p.Nome, &p.Descricao); erro != nil {
			respostaErro(w, http.StatusInternalServerError, "Erro ao ler produto", erro)
			return
		}
		produtos = append(produtos, p)
	}

	respostaJSON(w, http.StatusOK, produtos)
}

// GetProduto pega um produto pelo id passado no Param
func GetProduto(w http.ResponseWriter, r *http.Request) {
	id, erro := strconv.Atoi(mux.Vars(r)["idProduto"])

	if erro != nil {
		respostaErro(w, http.StatusBadRequest, "ID inválido", erro)
		return
	}

	banco, erro := database.Conectar()
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro ao conectar banco", erro)
		return
	}
	defer banco.Close()

	row := banco.QueryRow("SELECT idproduto, nome, descricao FROM produto WHERE idproduto = ?", id)

	var p produto
	if erro := row.Scan(&p.IDproduto, &p.Nome, &p.Descricao); erro != nil {
		respostaErro(w, http.StatusNotFound, "Produto não encontrado", erro)
		return
	}

	respostaJSON(w, http.StatusOK, p)
}

// AtualizarProduto atualiza um produto no banco de dados
func AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	id, erro := strconv.Atoi(mux.Vars(r)["idProduto"])
	if erro != nil {
		respostaErro(w, http.StatusBadRequest, "ID inválido", erro)
		return
	}

	var p produto
	if erro := json.NewDecoder(r.Body).Decode(&p); erro != nil {
		respostaErro(w, http.StatusBadRequest, "JSON inválido", erro)
		return
	}

	banco, erro := database.Conectar()
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Falha ao conectar banco", erro)
		return
	}
	defer banco.Close()

	stmt, erro := banco.Prepare("UPDATE produto SET nome = ?, descricao = ? WHERE idproduto = ?")
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro no prepare", erro)
		return
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(p.Nome, p.Descricao, id); erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro ao atualizar produto", erro)
		return
	}

	respostaJSON(w, http.StatusOK, map[string]string{"mensagem": "Produto atualizado"})
}

// DeletarProduto deleta um produto no banco de dados
func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	id, erro := strconv.Atoi(mux.Vars(r)["idProduto"])
	if erro != nil {
		respostaErro(w, http.StatusBadRequest, "ID inválido", erro)
		return
	}

	banco, erro := database.Conectar()
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro ao conectar banco", erro)
		return
	}
	defer banco.Close()

	stmt, erro := banco.Prepare("DELETE FROM produto WHERE idproduto = ?")
	if erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro no prepare", erro)
		return
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(id); erro != nil {
		respostaErro(w, http.StatusInternalServerError, "Erro ao deletar produto", erro)
		return
	}

	respostaJSON(w, http.StatusOK, map[string]string{"mensagem": "Produto deletado"})
}
