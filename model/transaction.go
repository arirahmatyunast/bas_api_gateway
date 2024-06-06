package model

import "time"

type Transaction struct {
	ID              int `gorm:"primarykey"`
	AccountID       string
	BankID          string
	Amount          int `gorm:"column:amount"`
	TransactionDate *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
