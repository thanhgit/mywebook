package api

import (
	"encoding/json"
	"github.com/imroc/req"
	"thanhgit.com/mywebhook/main/dto"
	"thanhgit.com/mywebhook/main/utils"
	variable2 "thanhgit.com/mywebhook/main/variable"
)

type AutoticketBot struct {
	Body dto.AutoticketBody
}

func (autoticket AutoticketBot)PushTicket() (dto.AutoticketResponse,dto.ResultDTO) {
	r := req.New()
	variable := variable2.GetInstance()
	url := variable.AutoticketURL

	header := req.Header{
		"Accept":	"application/json",
		"Authorization": variable2.GetInstance().AuthorizationAutoticket,
	}

	response, err := r.Post(url, header, req.BodyJSON(autoticket.Body))
	if err != nil {
		utils.PrintLog("Error at PushTicket function")
		panic(err)
	}

	result := dto.ResultDTO{}
	autoticketResponse := dto.AutoticketResponse{}
	if json.Unmarshal(response.Bytes(), &autoticketResponse) != nil {
		json.Unmarshal(response.Bytes(), &result)
	}

	return autoticketResponse, result
}
