package model

import "time"

type Transaction struct {
	ID              int        `gorm:"primarykey;autoIncrement" json:"id"`
	AccountID       string     `gorm:"foreignykey" json:"account_id"`
	BankID          string     `gorm:"foreignykey" json:"bank_id"`
	Amount          float64    `gorm:"column:amount" json:"amount" `
	TransactionDate *time.Time `gorm:"column:transaction_date" json:"transaction_date"`
}

func (a *Transaction) TableName() string {
	return "transaction"
}
