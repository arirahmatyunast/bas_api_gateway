package model

type Account struct {
	AccountID string `gorm:"primarykey" json:"account_id"`
	Name      string
	Username  string `gorm:"column-username"`
	Password  string
}

func (a *Account) TableName() string {
	return "account"
}
