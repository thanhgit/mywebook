package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type AOPTGetCommandBody struct {
	Token	string	`json:"token"`
	Id		int		`json:"id"`
}

func (aopt *AOPTGetCommandBody)ToString() string {
	result, err := json.Marshal(aopt)
	if err != nil {
		utils.PrintLog("Error at dto/AOPTGetCommandBody/ToString")
		panic(err)
	}

	return string(result)
}
