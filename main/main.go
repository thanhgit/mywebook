package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"thanhgit.com/mywebhook/main/api"
	"thanhgit.com/mywebhook/main/docs"
	_ "thanhgit.com/mywebhook/main/docs"
	"thanhgit.com/mywebhook/main/dto"
	"thanhgit.com/mywebhook/main/utils"
	"thanhgit.com/mywebhook/main/variable"
	"strconv"
	timelib "time"
)



func setupGlobalMiddleware(c *gin.Context) *gin.Context {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	return c
}

func pushAutoticket() (dto.AutoticketResponse, dto.ResultDTO) {
	autoticket := api.AutoticketBot{}

	_objectId := timelib.Now().Unix()%1000000
	_eventId := timelib.Now().Unix()%100000000

	autoticket.Body = dto.NewAutoticketBody(_objectId, _eventId)
	println(autoticket.Body.ToString())
	obj, err := autoticket.PushTicket()

	return obj, err
}

func pushAutoticketAutoCallFail() (dto.AutoticketResponse, dto.ResultDTO) {
	autoticket := api.AutoticketBot{}

	_objectId := timelib.Now().Unix()%1000000
	_eventId := timelib.Now().Unix()%100000000

	autoticket.Body = dto.NewAutoticketAutoCallFailBody(_objectId, _eventId)
	println(autoticket.Body.ToString())
	obj, err := autoticket.PushTicket()

	return obj, err
}


func pushAopt() string {
	aoptBot := api.AOPTBot{
		Username: variable.GetInstance().AOPTUsername,
		Password: variable.GetInstance().AOPTPassword,
	}

	runService112, err :=aoptBot.RunServiceBackup112()
	runService113, err :=aoptBot.RunServiceBackup113()
	if err.Code != 0 {
		return err.Result
	} else {
		jobId112 := runService112.Id
		jobId113 := runService113.Id
		timelib.Sleep(12* timelib.Second)
		getStatus112, err := aoptBot.GetStatus(jobId112)
		getStatus113, err := aoptBot.GetStatus(jobId113)


		if err.Code != 0 {
			return err.Result
		} else {
			return "[" + runService112.ToString() + " or " + getStatus112.ToString() + "][" + runService113.ToString() + " or " + getStatus113.ToString() + "]"
		}
	}
}

func main() {
	// Initialization global variable
	instanceVariable := variable.GetInstance()
	//loading config file
	config := new(utils.Config)
	err := config.Read("./main/webhook.yaml")
	if err == nil {
		config.ConvertToVariable(instanceVariable)
	}

	utils.PrintLog("Default configuration: \n" + instanceVariable.ToString())

	// initialization gin and swagger
	r := gin.New()
	docs.ReRegisterSwagger()
	url := ginSwagger.URL("http://"+ variable.GetInstance().IP+":"+strconv.Itoa(int(variable.GetInstance().Port))+"/swagger/doc.json") // The url point

	// Validate API and Input
	if !utils.IsValidateAPI() {
		return
	}

	// API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET(utils.ConvertAPI(variable.GetInstance().MY_ALERT_WELCOME), func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		utils.PrintLog("----------"+variable.GetInstance().MY_ALERT_WELCOME+"-------------------")
		c.JSON(200, gin.H{
			"Status": "Hi everyone! I am a my web hook",
			"Code": 200,
		})
	})

	r.POST(utils.ConvertAPI(variable.GetInstance().MY_ALERT_OUT_OF_CREDIT), func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		utils.PrintLog("----------"+variable.GetInstance().MY_ALERT_OUT_OF_CREDIT+"-------------------")
		autoTicket, err := pushAutoticket()
		str := pushAopt()

		if err.Code == 0 {
			c.JSON(500, gin.H{
				"Status": "Failled: " + autoTicket.ToString() + " or " + str,
				"Code": 200,
			})
		} else {
			c.JSON(200, gin.H{
				"Status": "Success:" + autoTicket.ToString() + " or " + str,
				"Code": 200,
			})
		}
	})

	r.POST(utils.ConvertAPI(variable.GetInstance().MY_ALERT_SERVER_FCALL_DOWN), func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		utils.PrintLog("----------"+variable.GetInstance().MY_ALERT_SERVER_FCALL_DOWN+"-------------------")
		autoTicket, err := pushAutoticket()
		str := pushAopt()

		if err.Code == 0 {
			c.JSON(200, gin.H{
				"Status": "Failled: " + err.Result + " or " + str,
				"Code": 500,
			})
		} else {
			c.JSON(200, gin.H{
				"Status": "Success:" + autoTicket.ToString() + " or " + str,
				"Code": 200,
			})
		}
	})

	r.POST(utils.ConvertAPI(variable.GetInstance().MY_ALERT_AUTO_CALL_FAIL), func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		utils.PrintLog("----------"+variable.GetInstance().MY_ALERT_AUTO_CALL_FAIL+"-------------------")
		autoTicket, err := pushAutoticketAutoCallFail()
		if err.Code == 0 {
			c.JSON(200, gin.H{
				"Status": "Failled: " + err.Result,
				"Code": 500,
			})
		} else {
			c.JSON(200, gin.H{
				"Status": "Success:" + autoTicket.ToString(),
				"Code": 200,
			})
		}
	})

	r.POST(utils.ConvertAPI(variable.GetInstance().MY_ALERT_SERVER_BACKUP_DOWN), func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		utils.PrintLog("----------"+variable.GetInstance().MY_ALERT_SERVER_BACKUP_DOWN+"-------------------")
		autoTicket, err := pushAutoticket()

		if err.Code == 0 {
			c.JSON(200, gin.H{
				"Status": "Failled: " + err.Result,
				"Code": 500,
			})
		} else {
			c.JSON(200, gin.H{
				"Status": "Success:" + autoTicket.ToString(),
				"Code": 200,
			})
		}
	})

	r.POST(utils.ConvertAPI(variable.GetInstance().MY_ALERT_TELEGRAM_SEND), func(c *gin.Context) {
		c = setupGlobalMiddleware(c)
		utils.PrintLog("----------"+variable.GetInstance().MY_ALERT_TELEGRAM_SEND+"-------------------")
		var notify dto.Notification

		bot := api.TelegramBot{}
		bot.GetInstance()

		if c.BindJSON(&notify) != nil {
			c.JSON(200, gin.H{
				"Status": "Error for posting server",
				"Code": 500,
			})
		} else {
			if notify.ChatId == 0 {
				if notify.Token == "" {
					bot.SendMessageToDefaultChannel(notify.Message)
				} else {
					bot.GetInstanceWithToken(notify.Token)
					bot.SendMessageToDefaultChannel(notify.Message)
				}

			} else {
				if notify.Token == "" {
					bot.SendMessage(notify.ChatId, notify.Message)
				} else {
					bot.GetInstanceWithToken(notify.Token)
					bot.SendMessage(notify.ChatId, notify.Message)
				}
			}

			c.JSON(200, gin.H{
				"Status": "Success",
				"Code": 200,
			})
		}
	})

	r.Run(variable.GetInstance().IP + ":" + strconv.Itoa(int(variable.GetInstance().Port)))
}


