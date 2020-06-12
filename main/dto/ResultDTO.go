package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type ResultDTO struct {
	Code	int		`json:"code"`
	Result	string	`json:"result"`
}

func (resultDTO *ResultDTO)ToString() string {
	result, err := json.Marshal(resultDTO)
	if err != nil {
		utils.PrintLog("Error at dto/ResultDTO/ToString")
		panic(err)
	}

	return string(result)
}