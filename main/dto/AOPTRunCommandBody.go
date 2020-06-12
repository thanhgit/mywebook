package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type AOPTRunCommandBody struct {
	Token 		string 		`json:token`
	DeviceIP	string		`json:device_ip`
	Command		string		`json:command`
}

func (aopt *AOPTRunCommandBody) ToString() string  {
	result, err := json.Marshal(aopt)
	if err != nil {
		utils.PrintLog("Error at dto/AOPTRunCommandBody/ToString")
		panic(err)
	}

	return string(result)
}