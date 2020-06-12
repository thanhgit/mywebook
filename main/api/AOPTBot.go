package api

import (
	"encoding/json"
	req "github.com/imroc/req"
	"thanhgit.com/mywebhook/main/dto"
	"thanhgit.com/mywebhook/main/utils"
	"thanhgit.com/mywebhook/main/variable"
	"strconv"
)

type AOPTBot struct {
	Token		string
	Username	string
	Password	string
}

func (aopt *AOPTBot) GetToken() string  {
	r := req.New()
	variable := variable.GetInstance()

	header := req.Header{
		"Accept":        "multipart/form-data",
	}

	tokenText, error := r.Post(variable.AoptURLGetToken, header, req.Param{"username": aopt.Username, "password": aopt.Password})
	if error != nil {
		panic(error)
	}
	token := dto.TokenDTO{}
	err := json.Unmarshal(tokenText.Bytes(), &token)

	if err != nil {
		utils.PrintLog("Error apt GetToken function")
		panic(err)
	}

	aopt.Token = token.TokenKey

	return token.TokenKey
}

func (aopt *AOPTBot) RunCommand(_deviceIp string, _command string) (dto.AOPTRunCommandResponse, dto.ResultDTO)  {
	r := req.New()
	variable := variable.GetInstance()
	header := req.Header{
		"Accept":        "application/json",
	}
	url := variable.AoptURLRunCommand
	token := aopt.GetToken()
	runCommand := dto.AOPTRunCommandBody{
		Token:    token,
		DeviceIP: _deviceIp,
		Command:  _command,
	}

	response, err := r.Post(url, header, req.BodyJSON(&runCommand))
	if err != nil {
		utils.PrintLog("Error at RunCommand function")
		panic(err)
	}

	runCommandResponse := dto.AOPTRunCommandResponse{}
	result := dto.ResultDTO{}
	if json.Unmarshal(response.Bytes(), &runCommandResponse) != nil {
		json.Unmarshal(response.Bytes(), &result)
	}

	return runCommandResponse, result
}

func (aopt *AOPTBot)GetCommand(_jobIp int)  (dto.AOPTGetCommandResponse, dto.ResultDTO){
	r := req.New()
	variable := variable.GetInstance()
	header := req.Header{
		"Accept":	"application/json",
	}
	url := variable.AoptURLGetCommand
	token := aopt.GetToken()
	getCommand := dto.AOPTGetCommandBody{
		Token: token,
		Id:    _jobIp,
	}

	response, err := r.Post(url, header, req.BodyJSON(&getCommand))
	if err != nil {
		utils.PrintLog("Error at GetCommand function")
		panic(err)
	}

	getCommandResponse := dto.AOPTGetCommandResponse{}
	result := dto.ResultDTO{}

	if json.Unmarshal(response.Bytes(), &getCommandResponse) != nil {
		json.Unmarshal(response.Bytes(), &result)
	}

	return getCommandResponse, result
}

func (aopt *AOPTBot)RunServiceBackup112() (dto.AOPTRunServiceResponse, dto.ResultDTO) {
	r := req.New()
	variable := variable.GetInstance()
	url := variable.AoptURLRunService
	header := req.Header{
		"Accept":	"application/json",
	}

	token := aopt.GetToken()
	_runService := dto.NewAOPTRunServiceBodyForBackup112(token)

	response, err := r.Post(url, header, req.BodyJSON(_runService))
	if err != nil {
		utils.PrintLog("Error at RunService function")
		panic(err)
	}

	runServiceResponse := dto.AOPTRunServiceResponse{}
	result := dto.ResultDTO{}
	if json.Unmarshal(response.Bytes(), &runServiceResponse) != nil {
		json.Unmarshal(response.Bytes(), &result)
	}

	return runServiceResponse, result
}

func (aopt *AOPTBot)RunServiceBackup113() (dto.AOPTRunServiceResponse, dto.ResultDTO) {
	r := req.New()
	variable := variable.GetInstance()
	url := variable.AoptURLRunService
	header := req.Header{
		"Accept":	"application/json",
	}

	token := aopt.GetToken()
	_runService := dto.NewAOPTRunServiceBodyForBackup113(token)

	response, err := r.Post(url, header, req.BodyJSON(_runService))
	if err != nil {
		utils.PrintLog("Error at RunService function")
		panic(err)
	}

	runServiceResponse := dto.AOPTRunServiceResponse{}
	result := dto.ResultDTO{}
	if json.Unmarshal(response.Bytes(), &runServiceResponse) != nil {
		json.Unmarshal(response.Bytes(), &result)
	}

	return runServiceResponse, result
}


func (aopt *AOPTBot)GetStatus(_jobId int64)  (dto.AOPTGetStatusResponse, dto.ResultDTO){
	r := req.New()
	variable := variable.GetInstance()
	url := variable.AoptURLGetStatus + strconv.Itoa(int(_jobId))
	header := req.Header{
		"Accept":	"application/json",
	}

	token := aopt.GetToken()
	getStatus := dto.NewAOPTGetStatusBodyForBackup(token)

	response, err := r.Post(url, header, req.BodyJSON(getStatus))
	if err != nil {
		utils.PrintLog("Error at GetStatus function")
		panic(err)
	}

	result := dto.ResultDTO{}
	getCommandResponse := dto.AOPTGetStatusResponse{}
	if json.Unmarshal(response.Bytes(), &getCommandResponse) != nil {
		json.Unmarshal(response.Bytes(), &result)
	}

	return getCommandResponse, result
}