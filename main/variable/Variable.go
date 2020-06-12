package variable

import (
	"strconv"
)


var instanceVariable *Variable

type Variable struct {
	// server
	IP 								string
	Port							int64
	Proxy 							string
	Debug							bool
	// telegram
	TelegramToken					string
	ChatID							int64
	// aopt
	AoptURLGetToken					string
	AoptURLRunCommand				string
	AoptURLGetCommand				string
	AoptURLRunService				string
	AoptURLGetStatus				string
	// autoticket
	AutoticketURL					string
	AuthorizationAutoticket			string
	// common
	VersionCode						int64
	AOPTUsername					string
	AOPTPassword					string
	// url
	MY_ALERT_WELCOME				string
	MY_ALERT_OUT_OF_CREDIT			string
	MY_ALERT_SERVER_FCALL_DOWN		string
	MY_ALERT_SERVER_BACKUP_DOWN	string
	MY_ALERT_AUTO_CALL_FAIL		string
	MY_ALERT_TELEGRAM_SEND			string

	// body for autoticket
	AutoticketBody struct{
		AlertPlan 				string	`yaml:"AlertPlan"`
		ServiceName				string	`yaml:"ServiceName"`
		ObjectID				int64	`yaml:"ObjectID"`
		KeyWord					string	`yaml:"KeyWord"`
		DedicateHostPlanning	string	`yaml:"DedicateHostPlanning"`
		EventID					int64	`yaml:"EventID"`
		Output					string	`yaml:"Output"`
		ServiceState			string	`yaml"ServiceState"`
		HostName				string	`yaml:"HostName"`
		HostState				string	`yaml:"HostState"`
		EventDateTime			string	`yaml:"EventDateTime"`
	}

	// body autoticket for autocall fail
	AutoticketAutoCallFailBody struct{
		AlertPlan 				string	`yaml:"AlertPlan"`
		ServiceName				string	`yaml:"ServiceName"`
		ObjectID				int64	`yaml:"ObjectID"`
		KeyWord					string	`yaml:"KeyWord"`
		DedicateHostPlanning	string	`yaml:"DedicateHostPlanning"`
		EventID					int64	`yaml:"EventID"`
		Output					string	`yaml:"Output"`
		ServiceState			string	`yaml"ServiceState"`
		HostName				string	`yaml:"HostName"`
		HostState				string	`yaml:"HostState"`
		EventDateTime			string	`yaml:"EventDateTime"`
	}

	AOPTRunServiceBody112 struct{
		Device			string 		`yaml:"Device"`
		Description		string		`yaml:"Description"`
		ServiceID		int64		`yaml:"ServiceID"`
		Token 			string		`yaml:"Token"`
		StringInterface	string		`yaml:"StringInterface"`
	}

	AOPTRunServiceBody113 struct{
		Device			string 		`yaml:"Device"`
		Description		string		`yaml:"Description"`
		ServiceID		int64		`yaml:"ServiceID"`
		Token 			string		`yaml:"Token"`
		StringInterface	string		`yaml:"StringInterface"`
	}
}

func GetInstance() *Variable {
	if  instanceVariable == nil {
		instanceVariable = new(Variable)
		instanceVariable.TelegramToken 							= ""
		instanceVariable.ChatID									= -1001467806884

		instanceVariable.AutoticketURL 							= "http://172.27.228.60:5664/nafautoticket/listalertinfo/"
		instanceVariable.AuthorizationAutoticket				= ""

		instanceVariable.AoptURLGetToken						= "https://tool.fpt.net/api/v1/login"
		instanceVariable.AoptURLRunCommand						= "https://tool.fpt.net/api/v1/run_command"
		instanceVariable.AoptURLGetCommand						= "https://tool.fpt.net/api/v1/get_command"
		instanceVariable.AoptURLRunService						= "https://tool.fpt.net/api/v1/jobs"
		instanceVariable.AoptURLGetStatus						= "https://tool.fpt.net/api/v1/job/"

		instanceVariable.VersionCode							= 6
		instanceVariable.AOPTUsername							= ""
		instanceVariable.AOPTPassword							= ""

		instanceVariable.IP										= "localhost"
		instanceVariable.Port									= 7000
		instanceVariable.Proxy									= ""
		instanceVariable.Debug									= true

		instanceVariable.MY_ALERT_WELCOME						= "/myalert"
		instanceVariable.MY_ALERT_OUT_OF_CREDIT				= "/myalert/outofcredit"
		instanceVariable.MY_ALERT_SERVER_FCALL_DOWN			= "/myalert/serverfcalldown"
		instanceVariable.MY_ALERT_SERVER_BACKUP_DOWN			= "/myalert/serverbackupdown"
		instanceVariable.MY_ALERT_AUTO_CALL_FAIL				= "/myalert/autocallfail"
		instanceVariable.MY_ALERT_TELEGRAM_SEND				= "/telegrams/send"

		// auto ticket
		instanceVariable.AutoticketBody.AlertPlan 				= ""
		instanceVariable.AutoticketBody.ServiceName 			= ""
		instanceVariable.AutoticketBody.ObjectID				= 0
		instanceVariable.AutoticketBody.KeyWord					= ""
		instanceVariable.AutoticketBody.DedicateHostPlanning	= ""
		instanceVariable.AutoticketBody.EventID					= 0
		instanceVariable.AutoticketBody.Output					= ""
		instanceVariable.AutoticketBody.ServiceState			= ""
		instanceVariable.AutoticketBody.HostName				= ""
		instanceVariable.AutoticketBody.HostState				= ""
		instanceVariable.AutoticketBody.EventDateTime			= ""

		// auto ticket for auto call faile
		instanceVariable.AutoticketAutoCallFailBody.AlertPlan 				= ""
		instanceVariable.AutoticketAutoCallFailBody.ServiceName 			= ""
		instanceVariable.AutoticketAutoCallFailBody.ObjectID				= 0
		instanceVariable.AutoticketAutoCallFailBody.KeyWord					= ""
		instanceVariable.AutoticketAutoCallFailBody.DedicateHostPlanning	= ""
		instanceVariable.AutoticketAutoCallFailBody.EventID					= 0
		instanceVariable.AutoticketAutoCallFailBody.Output					= ""
		instanceVariable.AutoticketAutoCallFailBody.ServiceState			= ""
		instanceVariable.AutoticketAutoCallFailBody.HostName				= ""
		instanceVariable.AutoticketAutoCallFailBody.HostState				= ""
		instanceVariable.AutoticketAutoCallFailBody.EventDateTime			= ""
		
		// aopt run service body 112
		instanceVariable.AOPTRunServiceBody112.Device			= "172.27.228.112"
		instanceVariable.AOPTRunServiceBody112.Description		= "Test service autocall failled 172.27.228.112"
		instanceVariable.AOPTRunServiceBody112.ServiceID		= 2636
		instanceVariable.AOPTRunServiceBody112.Token			= ""
		instanceVariable.AOPTRunServiceBody112.StringInterface	= ""

		// aopt run service body 113
		instanceVariable.AOPTRunServiceBody113.Device			= "172.27.228.113"
		instanceVariable.AOPTRunServiceBody113.Description		= "Test service autocall failled 172.27.228.113"
		instanceVariable.AOPTRunServiceBody113.ServiceID		= 2636
		instanceVariable.AOPTRunServiceBody113.Token			= ""
		instanceVariable.AOPTRunServiceBody113.StringInterface	= ""
	}

	return instanceVariable
}

func (variable *Variable)ToString() string  {
	result := ""
	splitStr := "\n"

	// server
	IPStr := "IP: " + variable.IP
	result += IPStr + splitStr

	PortStr := "Port: " + strconv.Itoa(int(variable.Port))
	result += PortStr + splitStr

	ProxyStr := "Proxy: " + variable.Proxy
	result += ProxyStr + splitStr

	debugStr := "Debug mode: " + strconv.FormatBool(variable.Debug)
	result += debugStr + splitStr

	result += splitStr

	// telegram
	telegramTokenStr := "Telegram Token: " + variable.TelegramToken
	result += telegramTokenStr + splitStr

	chatIDStr := "Chat ID: " + strconv.Itoa(int(variable.ChatID))
	result += chatIDStr + splitStr

	result += splitStr

	// autoticket
	autoticketURLStr := "Autoticket URL: " + variable.AutoticketURL
	result += autoticketURLStr + splitStr

	authorizationAutoticketStr := "Authorization Autoticket: " + variable.AuthorizationAutoticket
	result += authorizationAutoticketStr + splitStr

	result += splitStr

	// aopt
	aoptURLGetCommandStr := "AOPT Get Command URL: " + variable.AoptURLGetCommand
	result += aoptURLGetCommandStr + splitStr

	aoptURLRunCommandStr := "AOPT Run Command URL: " + variable.AoptURLRunCommand
	result += aoptURLRunCommandStr + splitStr

	aoptURLGetTokenStr := "AOPT Get Token URL: " + variable.AoptURLGetToken
	result += aoptURLGetTokenStr + splitStr

	aoptURLRunServiceStr := "AOPT Run Service URL: " + variable.AoptURLRunService
	result += aoptURLRunServiceStr + splitStr

	aoptURLGetStatusStr := "AOPT Get Status URL: " + variable.AoptURLGetStatus
	result += aoptURLGetStatusStr + splitStr

	result += splitStr

	// url
	url_my_alert_welcomeStr := "Welcome URL: " + variable.MY_ALERT_WELCOME
	result += url_my_alert_welcomeStr + splitStr

	url_my_alert_out_of_creditStr := "Out of credit URL: " + variable.MY_ALERT_OUT_OF_CREDIT
	result += url_my_alert_out_of_creditStr + splitStr

	url_my_alert_server_fcall_downStr := "Server Fcall down URL: " + variable.MY_ALERT_SERVER_FCALL_DOWN
	result += url_my_alert_server_fcall_downStr + splitStr

	url_my_alert_server_backup_downStr := "Server backup down URL: " + variable.MY_ALERT_SERVER_BACKUP_DOWN
	result += url_my_alert_server_backup_downStr + splitStr

	url_auto_call_faileStr := "Autocall fail URL: " + variable.MY_ALERT_AUTO_CALL_FAIL
	result += url_auto_call_faileStr + splitStr

	url_telegram_sendStr := "Telegram send URL: " + variable.MY_ALERT_TELEGRAM_SEND
	result += url_telegram_sendStr + splitStr

	result += splitStr

	// common
	versionCodeStr := "Version code: " + strconv.Itoa(int(variable.VersionCode))
	result += versionCodeStr + splitStr

	AOPTUsernameStr := "AOPT User name: " + variable.AOPTUsername
	result += AOPTUsernameStr + splitStr

	AOPTPasswordStr := "AOPT Password: " + variable.AOPTPassword
	result += AOPTPasswordStr + splitStr

	result += splitStr

	// auto ticket body
	AutoTicketStr := "---Auto Ticket---"
	result += AutoTicketStr + splitStr

	AutoTicket_AlertPlanStr := "Alert Plan: " + instanceVariable.AutoticketBody.AlertPlan
	result += AutoTicket_AlertPlanStr + splitStr

	AutoTicket_ServiceNameStr := "Service Name: " + instanceVariable.AutoticketBody.ServiceName
	result += AutoTicket_ServiceNameStr + splitStr

	AutoTicket_ObjectIDStr := "ObjectID: " + strconv.Itoa(int(instanceVariable.AutoticketBody.ObjectID))
	result += AutoTicket_ObjectIDStr + splitStr

	AutoTicket_KeywordStr := "Keyboard: " + instanceVariable.AutoticketBody.KeyWord
	result += AutoTicket_KeywordStr + splitStr

	AutoTicket_DedicatedHostPlanningStr := "Dedicated Host Planning: " + instanceVariable.AutoticketBody.DedicateHostPlanning
	result += AutoTicket_DedicatedHostPlanningStr + splitStr

	AutoTicket_EventIDStr := "Event ID: " + strconv.Itoa(int(instanceVariable.AutoticketBody.EventID))
	result += AutoTicket_EventIDStr + splitStr

	AutoTicket_OutputStr := "Output: " + instanceVariable.AutoticketBody.Output
	result += AutoTicket_OutputStr + splitStr

	AutoTicket_ServiceStateStr := "Service State: " + instanceVariable.AutoticketBody.ServiceState
	result += AutoTicket_ServiceStateStr + splitStr

	AutoTicket_HostNameStr := "Host Name: " + instanceVariable.AutoticketBody.HostName
	result += AutoTicket_HostNameStr + splitStr

	AutoTicket_HostStateStr := "Host State: " + instanceVariable.AutoticketBody.HostState
	result += AutoTicket_HostStateStr + splitStr

	AutoTicket_EventDateTimeStr := "Event Datetime: " + instanceVariable.AutoticketBody.EventDateTime
	result += AutoTicket_EventDateTimeStr + splitStr

	result += splitStr

	// auto ticket body for autcall fail
	AutoTicketAutoCallFailStr := "---Auto Ticket For autocall fail---"
	result += AutoTicketAutoCallFailStr + splitStr

	AutoTicketAutoCallFailAutoCallFail_AlertPlanStr := "Alert Plan: " + instanceVariable.AutoticketAutoCallFailBody.AlertPlan
	result += AutoTicketAutoCallFailAutoCallFail_AlertPlanStr + splitStr

	AutoTicketAutoCallFail_ServiceNameStr := "Service Name: " + instanceVariable.AutoticketAutoCallFailBody.ServiceName
	result += AutoTicketAutoCallFail_ServiceNameStr + splitStr

	AutoTicketAutoCallFail_ObjectIDStr := "ObjectID: " + strconv.Itoa(int(instanceVariable.AutoticketAutoCallFailBody.ObjectID))
	result += AutoTicketAutoCallFail_ObjectIDStr + splitStr

	AutoTicketAutoCallFail_KeywordStr := "Keyboard: " + instanceVariable.AutoticketAutoCallFailBody.KeyWord
	result += AutoTicketAutoCallFail_KeywordStr + splitStr

	AutoTicketAutoCallFail_DedicatedHostPlanningStr := "Dedicated Host Planning: " + instanceVariable.AutoticketAutoCallFailBody.DedicateHostPlanning
	result += AutoTicketAutoCallFail_DedicatedHostPlanningStr + splitStr

	AutoTicketAutoCallFail_EventIDStr := "Event ID: " + strconv.Itoa(int(instanceVariable.AutoticketAutoCallFailBody.EventID))
	result += AutoTicketAutoCallFail_EventIDStr + splitStr

	AutoTicketAutoCallFail_OutputStr := "Output: " + instanceVariable.AutoticketAutoCallFailBody.Output
	result += AutoTicketAutoCallFail_OutputStr + splitStr

	AutoTicketAutoCallFail_ServiceStateStr := "Service State: " + instanceVariable.AutoticketAutoCallFailBody.ServiceState
	result += AutoTicketAutoCallFail_ServiceStateStr + splitStr

	AutoTicketAutoCallFail_HostNameStr := "Host Name: " + instanceVariable.AutoticketAutoCallFailBody.HostName
	result += AutoTicketAutoCallFail_HostNameStr + splitStr

	AutoTicketAutoCallFail_HostStateStr := "Host State: " + instanceVariable.AutoticketAutoCallFailBody.HostState
	result += AutoTicketAutoCallFail_HostStateStr + splitStr

	AutoTicketAutoCallFail_EventDateTimeStr := "Event Datetime: " + instanceVariable.AutoticketAutoCallFailBody.EventDateTime
	result += AutoTicketAutoCallFail_EventDateTimeStr + splitStr

	result += splitStr


	// aopt run service body 112
	AOPTRunServiceBody112Str := "---AOPT Run Service Body 112---"
	result += AOPTRunServiceBody112Str + splitStr

	AOPTRunServiceBody112_DeviceStr := "Device: " + instanceVariable.AOPTRunServiceBody112.Device
	result += AOPTRunServiceBody112_DeviceStr + splitStr

	AOPTRunServiceBody112_DescriptionStr := "Description: " + instanceVariable.AOPTRunServiceBody112.Description
	result += AOPTRunServiceBody112_DescriptionStr + splitStr

	AOPTRunServiceBody112_ServiceIDStr := "Service ID: " + strconv.Itoa(int(instanceVariable.AOPTRunServiceBody112.ServiceID))
	result += AOPTRunServiceBody112_ServiceIDStr + splitStr

	AOPTRunServiceBody112_TokenStr := "Token: " + instanceVariable.AOPTRunServiceBody112.Token
	result += AOPTRunServiceBody112_TokenStr + splitStr

	AOPTRunServiceBody112_StringInterfaceStr := "String Interface: " + instanceVariable.AOPTRunServiceBody112.StringInterface
	result += AOPTRunServiceBody112_StringInterfaceStr + splitStr

	result += splitStr

	// aopt run service body 113
	AOPTRunServiceBody113Str := "---AOPT Run Service Body 113---"
	result += AOPTRunServiceBody113Str + splitStr

	AOPTRunServiceBody113_DeviceStr := "Device: " + instanceVariable.AOPTRunServiceBody113.Device
	result += AOPTRunServiceBody113_DeviceStr + splitStr

	AOPTRunServiceBody113_DescriptionStr := "Description: " + instanceVariable.AOPTRunServiceBody113.Description
	result += AOPTRunServiceBody113_DescriptionStr + splitStr

	AOPTRunServiceBody113_ServiceIDStr := "Service ID: " + strconv.Itoa(int(instanceVariable.AOPTRunServiceBody113.ServiceID))
	result += AOPTRunServiceBody113_ServiceIDStr + splitStr

	AOPTRunServiceBody113_TokenStr := "Token: " + instanceVariable.AOPTRunServiceBody113.Token
	result += AOPTRunServiceBody113_TokenStr + splitStr

	AOPTRunServiceBody113_StringInterfaceStr := "String Interface: " + instanceVariable.AOPTRunServiceBody113.StringInterface
	result += AOPTRunServiceBody113_StringInterfaceStr + splitStr

	return result
}
