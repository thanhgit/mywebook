package dto

import (
	"encoding/json"
	"thanhgit.com/mywebhook/main/utils"
)

type Notification struct {
	ChatId int64
	Message string `json:"text"`
	Token   string `json:"token"`
}

func (notify *Notification) ToString() string {
	result, err := json.Marshal(notify)
	if err != nil {
		utils.PrintLog("Error at dto/Notification/ToString")
		panic(err)
	}

	return string(result)
}