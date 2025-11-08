package model

import "time"

// Produto representa a tabela produto
type Produto struct {
	ID        uint32    `json:"idproduto,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Descricao string    `json:"descricao,omitempty"`
	IDUsuario uint32    `json:"idusuario,omitempty"`
	CriadoEm  time.Time `json:"criadoem"`
}
