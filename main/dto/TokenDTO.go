package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type TokenDTO struct {
	TokenKey string `json:"token_key"`
}

func (tokenDTO *TokenDTO)ToString() string {
	result, err := json.Marshal(tokenDTO)
	if err != nil {
		utils.PrintLog("Error at dto/TokenDTO/ToString")
		panic(err)
	}

	return string(result)
}