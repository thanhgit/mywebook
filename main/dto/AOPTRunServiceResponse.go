package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type AOPTRunServiceResponse struct {
	Code	int64	`json:"code"`
	Id 		int64 	`json:"id""`
	Result 	string	`json:"result"`
}

func (aopt *AOPTRunServiceResponse)ToString() string   {
	result, err := json.Marshal(aopt)
	if err != nil {
		utils.PrintLog("Error at dto/AOPTRunServiceResponse/ToString")
		panic(err)
	}

	return string(result)
}