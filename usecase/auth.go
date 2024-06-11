package usecase

import (
	"api_gateway/model"
	"api_gateway/utils"
)

//"api_gateway/handler"
//"github.com/gin-gonic/gin"

type Login struct {
}

type LoginInterface interface {
	Autentikasi(username, password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (a *Login) Autentikasi(username, password string) bool {

	accounts := model.Account{}
	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()
	orm.Find(&accounts, "username = ? AND password = ?", username, password)
	if accounts.AccountID == "" {
		return false
	}
	return true
	// if username == "admin" && password == "admin123" {
	// 	return true
	// }
	// return false
}
