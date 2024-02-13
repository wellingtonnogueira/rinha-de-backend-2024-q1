package domain

import (
	"time"
)

type AccountDetails struct {
	AccountBalance Balance `json:"balanco,omitempty"`
	LastTransaction []Transaction `json:"ultimas_transacoes,omitempty"`

}

type Balance struct {
	Total					uint		`json:"total,omitempty"`
	TransactionType			uint		`json:"limite,omitempty"`
	AccountStatementDate	time.Time	`json:"data_extrato"`
}