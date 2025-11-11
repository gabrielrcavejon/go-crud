package response

import (
	"encoding/json"
	"net/http"
)

// SucessoResponse é o padrao de retorno utilizado quando a requisição teve sucesso
type SucessoResponse struct {
	Sucesso  bool        `json:"sucesso"`
	Data     interface{} `json:"data,omitempty"`
	Mensagem string      `json:"mensagem,omitempty"`
}

// ErroResponse é o padrao de retorno utilizado quando a requisição nao teve sucesso
type ErroResponse struct {
	Sucesso bool        `json:"sucesso"`
	Erro    DetalheErro `json:"erro"`
}

// DetalheErro sao os detalhes do erro
type DetalheErro struct {
	Mensagem string `json:"mensagem"`
}

// RetonarSucesso é o retorno padrao da api em caso de sucesso
func RetonarSucesso(w http.ResponseWriter, status int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(SucessoResponse{
		Sucesso:  true,
		Data:     data,
		Mensagem: msg,
	})
}

// RetonarErro é o retorno padrao da api em caso de falha
func RetonarErro(w http.ResponseWriter, status int, mensagem string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErroResponse{
		Sucesso: false,
		Erro: DetalheErro{
			Mensagem: mensagem,
		},
	})
}
