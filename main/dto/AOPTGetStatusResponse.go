package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type StackObj struct {
	Result string	`json:"result"`
}

type TaskObj struct {
	Device 		string	`json:device`
	Id 			string	`json:"id"`
	Name 		string 	`json:"name"`
	Status 		string	`json:"status"`
	Task 		int 	`json:"task"`
	Template 	int 	`json:"template"`
}
type AOPTGetStatusResponse struct {
	FinishTime		string		`json:"finish_time"`
	ScheduleTime	string		`json:"schedule_time"`
	Stack 			StackObj	`json:"stack"`
	StartTime		string 		`json:start_time`
	Tasks			[]TaskObj	`json:tasks`
}



func (aopt *AOPTGetStatusResponse)ToString() string  {
	result, err := json.Marshal(aopt)
	if err != nil {
		utils.PrintLog("Error at dto/AOPTGetStatusResponse/ToString")
		panic(err)
	}

	return string(result)
}
