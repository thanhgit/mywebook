package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
	"thanhgit.com/mywebhook/main/variable"
)

type AOPTRunServiceBody struct {
	Device			string 		`json:"device"`
	Description		string		`json:"description"`
	ServiceId		int64		`json:"service_id"`
	Token 			string		`json:"token"`
	StringInterface	string		`json:"string_interface""`
}

func NewAOPTRunServiceBody(_device string, _description string, _serviceId int64, _token string, _stringInterface string) AOPTRunServiceBody {
	return AOPTRunServiceBody{
		_device,
		_description,
		_serviceId,
		_token,
		_stringInterface,
	}
}

func NewAOPTRunServiceBodyForBackup112(_token string) AOPTRunServiceBody {
	return NewAOPTRunServiceBody(
		variable.GetInstance().AOPTRunServiceBody112.Device,
		variable.GetInstance().AOPTRunServiceBody112.Description,
		variable.GetInstance().AOPTRunServiceBody112.ServiceID,
		_token,
		variable.GetInstance().AOPTRunServiceBody112.StringInterface,
		)
}

func NewAOPTRunServiceBodyForBackup113(_token string) AOPTRunServiceBody {
	return NewAOPTRunServiceBody(
		variable.GetInstance().AOPTRunServiceBody113.Device,
		variable.GetInstance().AOPTRunServiceBody113.Description,
		variable.GetInstance().AOPTRunServiceBody113.ServiceID,
		_token,
		variable.GetInstance().AOPTRunServiceBody113.StringInterface,
	)
}



func (aopt *AOPTRunServiceBody)ToString() string {
	result, err := json.Marshal(aopt)
	if err != nil {
		utils.PrintLog("Error at dto/AOPTRunServiceBody/ToString ")
		panic(err)
	}

	return string(result)
}