package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/web"
	"thanhgit.com/mywebhook/main/utils"
	"strconv"
)

//
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c *gin.Context) {
	err := web.APIError{}
	utils.PrintLog(strconv.Itoa(err.ErrorCode) + " " + err.CreatedAt.String() + " " +err.ErrorMessage)
}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(c *gin.Context) {
	utils.PrintLog("RESTful API")
}