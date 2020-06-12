package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type AOPTGetCommandResponse struct {
	Output	[]interface{}	`json:"output"`
	Result	bool			`json:"result"`
	State 	string			`json:"state"`
}

func (aopt AOPTGetCommandResponse)ToString() string {
	result, err :=json.Marshal(&aopt)
	if err != nil {
		utils.PrintLog("Error dto/AOPTGETCommandResponse/ToString")
		panic(err)
	}

	return string(result)
}