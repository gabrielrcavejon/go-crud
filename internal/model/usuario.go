package model

import "time"

// Usuario representa a tabela usuario
type Usuario struct {
	ID       uint32    `json:"idusuario,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem"`
}
