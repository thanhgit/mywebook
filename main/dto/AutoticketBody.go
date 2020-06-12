package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
	"thanhgit.com/mywebhook/main/variable"
)

type AutoticketBody struct {
	AlertPlan 				string	`json:"AlertPlan"`
	ServiceName				string	`json:"ServiceName"`
	ObjectID				int64	`json:"ObjectID"`
	KeyWord					string	`json:"KeyWord"`
	DedicateHostPlanning	string	`json:"DedicateHostPlanning"`
	EventID					int64	`json:"EventID"`
	Output					string	`json:"Output""`
	ServiceState			string	`json"ServiceState"`
	HostName				string	`json:"HostName"`
	HostState				string	`json:"HostState"`
	EventDateTime			string	`json:"EventDateTime"`
}

func (autoticket *AutoticketBody) ToString() string {
	result, err := json.Marshal(autoticket)
	if err != nil {
		utils.PrintLog("Error at ToString of AutoticketBody")
		panic(err)
	}

	return string(result)
}

func NewAutoticketBody(_objectID int64, _eventID int64) AutoticketBody {
	time := utils.GetInstance()
	autoticketBody := AutoticketBody{
		AlertPlan:            variable.GetInstance().AutoticketBody.AlertPlan,
		ServiceName:          variable.GetInstance().AutoticketBody.ServiceName,
		ObjectID:             _objectID,
		KeyWord:              variable.GetInstance().AutoticketBody.KeyWord,
		DedicateHostPlanning: variable.GetInstance().AutoticketBody.DedicateHostPlanning,
		EventID:              _eventID,
		Output:               variable.GetInstance().AutoticketBody.Output,
		ServiceState:         variable.GetInstance().AutoticketBody.ServiceState,
		HostName:             variable.GetInstance().AutoticketBody.HostName,
		HostState:            variable.GetInstance().AutoticketBody.HostState,
		EventDateTime:        time.GetNowWithFormatOfAutoticket(),
	}

	return autoticketBody
}

func NewAutoticketAutoCallFailBody(_objectID int64, _eventID int64) AutoticketBody {
	time := utils.GetInstance()
	autoticketBody := AutoticketBody{
		AlertPlan:            variable.GetInstance().AutoticketAutoCallFailBody.AlertPlan,
		ServiceName:          variable.GetInstance().AutoticketAutoCallFailBody.ServiceName,
		ObjectID:             _objectID,
		KeyWord:              variable.GetInstance().AutoticketAutoCallFailBody.KeyWord,
		DedicateHostPlanning: variable.GetInstance().AutoticketAutoCallFailBody.DedicateHostPlanning,
		EventID:              _eventID,
		Output:               variable.GetInstance().AutoticketAutoCallFailBody.Output,
		ServiceState:         variable.GetInstance().AutoticketAutoCallFailBody.ServiceState,
		HostName:             variable.GetInstance().AutoticketAutoCallFailBody.HostName,
		HostState:            variable.GetInstance().AutoticketAutoCallFailBody.HostState,
		EventDateTime:        time.GetNowWithFormatOfAutoticket(),
	}

	return autoticketBody
}