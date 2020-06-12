package utils

import (
	"github.com/gin-gonic/gin"
	"thanhgit.com/mywebhook/main/variable"
	"strings"
)

/*
* Input: 	string /telegram/send
* Output: 	string /api/telegram/send
 */
func ConvertAPI(uri string) string  {
	return "/api" + uri
}

/*
* Input:
* 	Status 	string
*	Code	int64
* 	c		*gin.Context
* Output:
 */
func ResponseAPI(_status string, _code int64, c *gin.Context)  {
	c.JSON(200, gin.H{
		"Status": _status,
		"Code": _code,
	})
}

/**

 */
func IsValidateAPI() bool {
	if variable.GetInstance().IP == "0.0.0.0" || strings.TrimSpace(variable.GetInstance().IP) == "" {
		PrintLog("Let reset your IP in webhook.yml with format 192.x.x.x")
		return false
	}

	return true
}

func PrintLog(_msg string) {
	if variable.GetInstance().Debug {
		println(_msg)
	}
}