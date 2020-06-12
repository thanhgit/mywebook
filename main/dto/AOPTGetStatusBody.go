package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type AOPTGetStatusBody struct {
	Token	string			`json:"token"`
	Slack	[]string		`json:"stack"`
}

func NewAOPTGetStatusBodyForBackup(_token string) AOPTGetStatusBody {
	var arrStr []string
	arrStr = append(arrStr, "result")
	return AOPTGetStatusBody{
		Token: _token,
		Slack: arrStr,
	}
}

func (aopt *AOPTGetStatusBody)ToString() string  {
	result, err := json.Marshal(aopt)
	if err != nil {
		utils.PrintLog("Error at dto/AOPTGetStatusBody/ToString")
		panic(err)
	}

	return string(result)
}