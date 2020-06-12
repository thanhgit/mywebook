package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type AutoticketResponse struct {
	ID						int64		`json:"id"`
	AlertPlan				string		`json:"AlertPlan"`
	ServiceName				string		`json:"ServiceName"`
	ObjectID				int64		`json:"ObjectID"`
	KeyWord					string		`json:"KeyWord"`
	DedicateHostPlanning	string		`json:"DedicateHostPlanning"`
	EventID					int64		`json:"EventID"`
	EventType				string		`json:"EventType`
	Output					string		`json:"Output""`
	ServiceState			string		`json"ServiceState"`
	HostName				string		`json:"HostName"`
	HostState				string		`json:"HostState"`
	EventDateTime			string		`json:"EventDateTime"`
	ReceiveDateTime			string		`json:"ReceiveDateTime"`
	isDownTime				bool		`json:"isDownTime"`
}

func (autoticket *AutoticketResponse)ToString() string  {
	result, err := json.Marshal(autoticket)
	if err != nil {
		utils.PrintLog("Error at dto/AutoticketResponse/ToString")
		panic(err)
	}

	return string(result)
}