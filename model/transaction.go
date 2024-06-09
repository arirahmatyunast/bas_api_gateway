package model

import "time"

type Transaction struct {
	ID              int `gorm:"primarykey"`
	AccountID       string
	BankID          string
	Amount          int        `gorm:"column:amount" `
	TransactionDate *time.Time `gorm:"column:transaction_date;not null" json:"transaction_date"`
}

func (a *Transaction) TableName() string {
	return "transaction"
}
