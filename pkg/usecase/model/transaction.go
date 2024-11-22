package model

import "github.com/kenta0518/mf-codetest/pkg/domain/entity"

type Transaction struct {
	UserID      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func NewTransaction(transaction *entity.Transaction) *Transaction {
	return &Transaction{
		UserID:      transaction.UserID,
		Amount:      transaction.Amount,
		Description: transaction.Description,
	}
}
