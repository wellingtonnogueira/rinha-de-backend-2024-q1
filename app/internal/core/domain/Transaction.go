package domain

import "time"

type Transaction struct {
	Value			uint	`json:"valor,omitempty"`
	TransactionType	string	`json:"tipo,omitempty"`
	Description		string	`json:"descricao,omitempty"`
	CreatedAt		time.Time `json:"realizada_em"`
}

type TransactionResponse struct {
	AccountLimit	uint	`json:"limite,omitempty"`
	AccountBalance	uint	`json:"saldo,omitempty"`
}