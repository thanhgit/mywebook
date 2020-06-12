package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type Account struct {
	Username 	string 	`json:"username"`
	Password	string	`json:"password"`
}

func (account *Account)ToString() string {
	result, err := json.Marshal(account)
	if err != nil {
		utils.PrintLog("Error at dto/Account/ToString")
		panic(err)
	}

	return string(result)
}