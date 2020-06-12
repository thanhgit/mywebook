package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type AOPTRunCommandResponse struct {
	Id 		string 		`json:"id"`
	Result	bool	 	`json:"result"`
}

func (aopt *AOPTRunCommandResponse) ToString() string  {
	result, err := json.Marshal(aopt)
	if err != nil {
		utils.PrintLog("Error at dto/AOPTRunCommandResponse/ToString")
		panic(err)
	}

	return string(result)
}
