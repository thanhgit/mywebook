package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"thanhgit.com/mywebhook/main/variable"
)

type IConfig interface {
	// Server
	SetIP(string)
	SetPort(int64)
	IsDebug(bool)
	// telegram
	SetTelegramToken(string)
	SetChatID(int64)
	// aopt
	SetAoptURLGetToken(string)
	SetAoptURLRunCommand(string)
	SetAoptURLGetCommand(string)
	SetAoptURLRunService(string)
	SetAoptURLGetStatus(string)
	// autoticket
	SetAutoticketURL(string)
	SetAuthorizationAutoticket(string)
	// common
	SetVersionCode(int64)
	SetAOPTUsername(string)
	SetAOPTPassword(string)
	// url
	SetMY_ALERT_WELCOME(string)
	SetMY_ALERT_OUT_OF_CREDIT(string)
	SetMY_ALERT_SERVER_FCALL_DOWN(string)
	SetMY_ALERT_SERVER_BACKUP_DOWN(string)
	SetMY_ALERT_AUTO_CALL_FAIL(string)
	SetMY_ALERT_TELEGRAM_SEND(string)

	// autoticket
	SetAutoticketBody(
		AlertPlan string,
		ServiceName				string,
		ObjectID				int64,
		KeyWord					string,
		DedicateHostPlanning	string,
		EventID					int64,
		Output					string,
		ServiceState			string,
		HostName				string,
		HostState				string,
		EventDateTime			string)

	// autoticket autocall fail
	SetAutoticketAutoCallFailBody(
		AlertPlan string,
		ServiceName				string,
		ObjectID				int64,
		KeyWord					string,
		DedicateHostPlanning	string,
		EventID					int64,
		Output					string,
		ServiceState			string,
		HostName				string,
		HostState				string,
		EventDateTime			string)

	// aopt run service body 112
	SetAOPTRunServiceBody112(
		Device			string,
		Description		string,
		ServiceID		int64,
		Token 			string,
		StringInterface	string)

	// aopt run service body 113
	SetAOPTRunServiceBody113(
		Device			string,
		Description		string,
		ServiceID		int64,
		Token 			string,
		StringInterface	string)

}

type AutoticketBodyYaml struct {
	AlertPlan 				string	`yaml:"AlertPlan"`
	ServiceName				string	`yaml:"ServiceName"`
	ObjectID				int64	`yaml:"ObjectID"`
	KeyWord					string	`yaml:"KeyWord"`
	DedicateHostPlanning	string	`yaml:"DedicateHostPlanning"`
	EventID					int64	`yaml:"EventID"`
	Output					string	`yaml:"Output"`
	ServiceState			string	`yaml:"ServiceState"`
	HostName				string	`yaml:"HostName"`
	HostState				string	`yaml:"HostState"`
	EventDateTime			string	`yaml:"EventDateTime"`
}

type AutoticketAutoCalFailBodyYaml struct {
	AlertPlan 				string	`yaml:"AlertPlan"`
	ServiceName				string	`yaml:"ServiceName"`
	ObjectID				int64	`yaml:"ObjectID"`
	KeyWord					string	`yaml:"KeyWord"`
	DedicateHostPlanning	string	`yaml:"DedicateHostPlanning"`
	EventID					int64	`yaml:"EventID"`
	Output					string	`yaml:"Output"`
	ServiceState			string	`yaml:"ServiceState"`
	HostName				string	`yaml:"HostName"`
	HostState				string	`yaml:"HostState"`
	EventDateTime			string	`yaml:"EventDateTime"`
}

type AOPTRunServiceBody112Yaml struct{
	Device			string 		`yaml:"Device"`
	Description		string		`yaml:"Description"`
	ServiceID		int64		`yaml:"ServiceID"`
	Token 			string		`yaml:"Token"`
	StringInterface	string		`yaml:"StringInterface"`
}

type AOPTRunServiceBody113Yaml struct{
	Device			string 		`yaml:"Device"`
	Description		string		`yaml:"Description"`
	ServiceID		int64		`yaml:"ServiceID"`
	Token 			string		`yaml:"Token"`
	StringInterface	string		`yaml:"StringInterface"`
}

type Config struct {
	// server
	IP 								string 		`yaml:"IP"`
	Port							int64		`yaml:"Port"`
	Proxy							string		`yaml:"Proxy"`
	Debug							bool		`yaml:"Debug"`
	// telegram
	TelegramToken					string  	`yaml:"TelegramToken"`
	ChatId							int64		`yaml:"ChatID"`
	// aopt
	AoptURLGetToken					string 		`yaml:"AoptURLGetToken"`
	AoptURLRunCommand				string 		`yaml:"AoptURLRunCommand"`
	AoptURLGetCommand				string 		`yaml:"AoptURLGetCommand"`
	AoptURLRunService				string 		`yaml:"AoptURLRunService"`
	AoptURLGetStatus				string 		`yaml:"AoptURLGetStatus"`
	// autoticket
	AutoticketURL					string  	`yaml:"AutoticketURL"`
	AuthorizationAutoticket			string		`yaml:"AuthorizationAutoticket"`
	// common
	VersionCode						int64 		`yaml:"VersionCode"`
	AOPTUsername					string 		`yaml:"AOPTUsername"`
	AOPTPassword					string 		`yaml:"AOPTPassword"`
	// url
	MY_ALERT_WELCOME				string		`yaml:"MY_ALERT_WELCOME"`
	MY_ALERT_OUT_OF_CREDIT			string		`yaml:"MY_ALERT_OUT_OF_CREDIT"`
	MY_ALERT_SERVER_FCALL_DOWN		string 		`yaml:"MY_ALERT_SERVER_FCALL_DOWN"`
	MY_ALERT_SERVER_BACKUP_DOWN	string 		`yaml:"MY_ALERT_SERVER_BACKUP_DOWN"`
	MY_ALERT_AUTO_CALL_FAIL		string		`yaml:"MY_ALERT_AUTO_CALL_FAIL"`
	MY_ALERT_TELEGRAM_SEND			string		`yaml:"MY_ALERT_TELEGRAM_SEND"`

	// auto ticket
	AutoticketBody 	AutoticketBodyYaml							`yaml:"AutoticketBody"`
	AutoticketAutoCallFailBody 	AutoticketAutoCalFailBodyYaml	`yaml:"AutoticketAutoCalFailBody"`

	// aopt run service body 112
	AOPTRunServiceBody112 AOPTRunServiceBody112Yaml	`yaml:"AOPTRunServiceBody112"`

	// aopt run service body 113
	AOPTRunServiceBody113 AOPTRunServiceBody113Yaml	`yaml:"AOPTRunServiceBody113"`
}

func (config *Config) Init()  {

}

type ConfigBuilder struct {
	iConfig *IConfig
	config  *Config
}

func (builder *ConfigBuilder) Init()  {
	builder.iConfig = new(IConfig)
	builder.config 	= new(Config)
}
 // region builder function
 // server
func (builder *ConfigBuilder) SetIP(_ip string) *IConfig {
	builder.config.IP = _ip
	return builder.iConfig
}

func (builder *ConfigBuilder) SetPort(_port int64) *IConfig {
	builder.config.Port = _port
	return builder.iConfig
}

func (builder *ConfigBuilder) SetProxy(_proxy string) *IConfig {
	builder.config.Proxy = _proxy
	return builder.iConfig
}

func (buidler *ConfigBuilder) IsDebug(_isDebug bool) *IConfig  {
	buidler.config.Debug = _isDebug
	return buidler.iConfig
}

// telegram
func (buider *ConfigBuilder) SetTelegramToken(_telegramToken string) *IConfig {
	buider.config.TelegramToken = _telegramToken
	return buider.iConfig
}

func (builder *ConfigBuilder) SetChatID(_chatID int64) *IConfig {
	builder.config.ChatId = _chatID
	return builder.iConfig
}

// aopt
func (builder *ConfigBuilder) SetAoptURLGetToken(_aoptURLGetToken string) *IConfig {
	builder.config.AoptURLGetToken = _aoptURLGetToken
	return builder.iConfig
}

func (builder *ConfigBuilder) SetAoptURLRunCommand(_aoptURLRunCommand string) *IConfig {
	builder.config.AoptURLRunCommand = _aoptURLRunCommand
	return builder.iConfig
}

func (builder *ConfigBuilder) SetAoptURLGetCommand(_aoptURLGetCommand string) *IConfig {
	builder.config.AoptURLGetCommand = _aoptURLGetCommand
	return builder.iConfig
}

func (builder *ConfigBuilder) SetAoptURLRunService(_aoptURLRunService string) *IConfig {
	builder.config.AoptURLRunService = _aoptURLRunService
	return builder.iConfig
}

func (builder *ConfigBuilder) SetAoptURLGetStatus(_aoptULRGetStatus string) *IConfig {
	builder.config.AoptURLGetStatus = _aoptULRGetStatus
	return builder.iConfig
}

// autoticket
func (builder *ConfigBuilder) SetAutoticketURL(_autoticketURL string) *IConfig {
	builder.config.AutoticketURL = _autoticketURL
	return builder.iConfig
}


func (builder *ConfigBuilder) SetAuthorizationAutoticket(_authorizationAutoticket string) *IConfig  {
	builder.config.AuthorizationAutoticket = _authorizationAutoticket
	return builder.iConfig
}
// common
func (builder *ConfigBuilder) SetVersioCode(_versionCode int64) *IConfig {
	builder.config.VersionCode = _versionCode
	return builder.iConfig
}

func (builder *ConfigBuilder) SetAOPTUsername(_username string) *IConfig {
	builder.config.AOPTUsername = _username
	return builder.iConfig
}

func (builder *ConfigBuilder) SetAOPTPassword(_password string) *IConfig  {
	builder.config.AOPTPassword = _password
	return builder.iConfig
}

// url
func (builder *ConfigBuilder) SetMY_ALERT_WELCOME(_my_alert_welcome string)	*IConfig {
	builder.config.MY_ALERT_WELCOME = _my_alert_welcome
	return builder.iConfig
}

func (builder *ConfigBuilder) SetMY_ALERT_OUT_OF_CREDIT(_my_alert_out_of_credit string) *IConfig {
	builder.config.MY_ALERT_OUT_OF_CREDIT = _my_alert_out_of_credit
	return builder.iConfig
}

func (builder *ConfigBuilder) SetMY_ALERT_SERVER_FCALL_DOWN(_my_alert_server_fcall_down string) *IConfig {
	builder.config.MY_ALERT_SERVER_FCALL_DOWN = _my_alert_server_fcall_down
	return builder.iConfig
}

func (builder *ConfigBuilder) SetMY_ALERT_SERVER_BACKUP_DOWN(_my_alert_server_backup_down string) *IConfig {
	builder.config.MY_ALERT_SERVER_BACKUP_DOWN = _my_alert_server_backup_down
	return builder.iConfig
}

func (builder *ConfigBuilder) SetMY_ALERT_AUTO_CALL_FAIL(_my_alert_auto_call_fail string) *IConfig {
	builder.config.MY_ALERT_AUTO_CALL_FAIL = _my_alert_auto_call_fail
	return builder.iConfig
}

func (builder *ConfigBuilder) SetMY_ALERT_TELEGRAM_SEND(_my_alert_telegram_send string) *IConfig {
	builder.config.MY_ALERT_TELEGRAM_SEND = _my_alert_telegram_send
	return builder.iConfig
}

// autoticket
func (builder *ConfigBuilder) SetAutoticketBody(
	AlertPlan string,
	ServiceName				string,
	ObjectID				int64,
	KeyWord					string,
	DedicateHostPlanning	string,
	EventID					int64,
	Output					string,
	ServiceState			string,
	HostName				string,
	HostState				string,
	EventDateTime			string) *IConfig {
	builder.config.AutoticketBody.AlertPlan 			= AlertPlan
	builder.config.AutoticketBody.ServiceName 			= ServiceName
	builder.config.AutoticketBody.ObjectID				= ObjectID
	builder.config.AutoticketBody.KeyWord				= KeyWord
	builder.config.AutoticketBody.DedicateHostPlanning	= DedicateHostPlanning
	builder.config.AutoticketBody.EventID				= EventID
	builder.config.AutoticketBody.Output				= Output
	builder.config.AutoticketBody.ServiceState			= ServiceState
	builder.config.AutoticketBody.HostName				= HostName
	builder.config.AutoticketBody.HostState				= HostState
	builder.config.AutoticketBody.EventDateTime			= EventDateTime
	return builder.iConfig
}

// autoticket
func (builder *ConfigBuilder) SetAutoticketAutoCallFailBody(
	AlertPlan string,
	ServiceName				string,
	ObjectID				int64,
	KeyWord					string,
	DedicateHostPlanning	string,
	EventID					int64,
	Output					string,
	ServiceState			string,
	HostName				string,
	HostState				string,
	EventDateTime			string) *IConfig {
	builder.config.AutoticketAutoCallFailBody.AlertPlan 			= AlertPlan
	builder.config.AutoticketAutoCallFailBody.ServiceName 			= ServiceName
	builder.config.AutoticketAutoCallFailBody.ObjectID				= ObjectID
	builder.config.AutoticketAutoCallFailBody.KeyWord				= KeyWord
	builder.config.AutoticketAutoCallFailBody.DedicateHostPlanning	= DedicateHostPlanning
	builder.config.AutoticketAutoCallFailBody.EventID				= EventID
	builder.config.AutoticketAutoCallFailBody.Output				= Output
	builder.config.AutoticketAutoCallFailBody.ServiceState			= ServiceState
	builder.config.AutoticketAutoCallFailBody.HostName				= HostName
	builder.config.AutoticketAutoCallFailBody.HostState				= HostState
	builder.config.AutoticketAutoCallFailBody.EventDateTime			= EventDateTime
	return builder.iConfig
}

// aopt run service body 112
func (builder *ConfigBuilder) SetAOPTRunServiceBody112(
	Device			string,
	Description		string,
	ServiceID		int64,
	Token 			string,
	StringInterface	string) *IConfig {
		builder.config.AOPTRunServiceBody112.Device 			= Device
		builder.config.AOPTRunServiceBody112.Description 		= Description
		builder.config.AOPTRunServiceBody112.ServiceID 			= ServiceID
		builder.config.AOPTRunServiceBody112.Token				= Token
		builder.config.AOPTRunServiceBody112.StringInterface	= StringInterface
		return builder.iConfig
}

// aopt run service body 113
func (builder *ConfigBuilder) SetAOPTRunServiceBody113(
	Device			string,
	Description		string,
	ServiceID		int64,
	Token 			string,
	StringInterface	string) *IConfig {
		builder.config.AOPTRunServiceBody113.Device 			= Device
		builder.config.AOPTRunServiceBody113.Description		= Description
		builder.config.AOPTRunServiceBody113.ServiceID			= ServiceID
		builder.config.AOPTRunServiceBody113.Token				= Token
		builder.config.AOPTRunServiceBody113.StringInterface	= StringInterface
		return builder.iConfig
}

func (builder *ConfigBuilder) Create() *Config {
	return builder.config
}

 // end region builder function

func (config *Config) Read(_path string) error  {
	data, err := ioutil.ReadFile(_path)
	if err != nil {
		PrintLog("Error at utils/FileUtils/Read")
		panic(err)
		return err
	}

	//PrintLog(string(data))

	errUnmarshal := yaml.Unmarshal([]byte(data), &config)

	if errUnmarshal != nil {
		PrintLog("Error at utils/FileUtils/Read/errUnmarshal")
		panic(err)
	}

	return nil
}

func (config *Config) ConvertToVariable(variable * variable.Variable)  {
	// server
	variable.IP 	= config.IP
	variable.Port 	= config.Port
	variable.Proxy  = config.Proxy
	variable.Debug 	= config.Debug

	// telegram
	variable.TelegramToken 	= config.TelegramToken
	variable.ChatID			= config.ChatId

	// aopt
	variable.AoptURLGetToken 	= config.AoptURLGetToken
	variable.AoptURLGetStatus 	= config.AoptURLGetStatus
	variable.AoptURLRunCommand 	= config.AoptURLRunCommand
	variable.AoptURLGetCommand 	= config.AoptURLGetCommand
	variable.AoptURLRunService 	= config.AoptURLRunService

	// autoticket
	variable.AutoticketURL 				= config.AutoticketURL
	variable.AuthorizationAutoticket 	= config.AuthorizationAutoticket

	// common
	variable.VersionCode 	= config.VersionCode
	variable.AOPTUsername  	= config.AOPTUsername
	variable.AOPTPassword  	= config.AOPTPassword

	// url
	variable.MY_ALERT_WELCOME 				= config.MY_ALERT_WELCOME
	variable.MY_ALERT_TELEGRAM_SEND 		= config.MY_ALERT_TELEGRAM_SEND
	variable.MY_ALERT_SERVER_BACKUP_DOWN 	= config.MY_ALERT_SERVER_BACKUP_DOWN
	variable.MY_ALERT_SERVER_FCALL_DOWN 	= config.MY_ALERT_SERVER_FCALL_DOWN
	variable.MY_ALERT_OUT_OF_CREDIT 		= config.MY_ALERT_OUT_OF_CREDIT
	variable.MY_ALERT_AUTO_CALL_FAIL		= config.MY_ALERT_AUTO_CALL_FAIL

	// auto ticket body
	variable.AutoticketBody.AlertPlan				= config.AutoticketBody.AlertPlan
	variable.AutoticketBody.ServiceName				= config.AutoticketBody.ServiceName
	variable.AutoticketBody.ObjectID				= config.AutoticketBody.ObjectID
	variable.AutoticketBody.KeyWord					= config.AutoticketBody.KeyWord
	variable.AutoticketBody.DedicateHostPlanning	= config.AutoticketBody.DedicateHostPlanning
	variable.AutoticketBody.EventID					= config.AutoticketBody.EventID
	variable.AutoticketBody.Output					= config.AutoticketBody.Output
	variable.AutoticketBody.ServiceState			= config.AutoticketBody.ServiceState
	variable.AutoticketBody.HostName				= config.AutoticketBody.HostName
	variable.AutoticketBody.HostState				= config.AutoticketBody.HostState
	variable.AutoticketBody.EventDateTime			= config.AutoticketBody.EventDateTime

	// auto ticket auto call fail body
	variable.AutoticketAutoCallFailBody.AlertPlan				= config.AutoticketAutoCallFailBody.AlertPlan
	variable.AutoticketAutoCallFailBody.ServiceName				= config.AutoticketAutoCallFailBody.ServiceName
	variable.AutoticketAutoCallFailBody.ObjectID				= config.AutoticketAutoCallFailBody.ObjectID
	variable.AutoticketAutoCallFailBody.KeyWord					= config.AutoticketAutoCallFailBody.KeyWord
	variable.AutoticketAutoCallFailBody.DedicateHostPlanning	= config.AutoticketAutoCallFailBody.DedicateHostPlanning
	variable.AutoticketAutoCallFailBody.EventID					= config.AutoticketAutoCallFailBody.EventID
	variable.AutoticketAutoCallFailBody.Output					= config.AutoticketAutoCallFailBody.Output
	variable.AutoticketAutoCallFailBody.ServiceState			= config.AutoticketAutoCallFailBody.ServiceState
	variable.AutoticketAutoCallFailBody.HostName				= config.AutoticketAutoCallFailBody.HostName
	variable.AutoticketAutoCallFailBody.HostState				= config.AutoticketAutoCallFailBody.HostState
	variable.AutoticketAutoCallFailBody.EventDateTime			= config.AutoticketAutoCallFailBody.EventDateTime

	// aopt run service body 112
	variable.AOPTRunServiceBody112.Device			= config.AOPTRunServiceBody112.Device
	variable.AOPTRunServiceBody112.Description		= config.AOPTRunServiceBody112.Description
	variable.AOPTRunServiceBody112.ServiceID		= config.AOPTRunServiceBody112.ServiceID
	variable.AOPTRunServiceBody112.Token			= config.AOPTRunServiceBody112.Token
	variable.AOPTRunServiceBody112.StringInterface	= config.AOPTRunServiceBody112.StringInterface

	// aopt run service body 113
	variable.AOPTRunServiceBody113.Device 			= config.AOPTRunServiceBody113.Device
	variable.AOPTRunServiceBody113.Description		= config.AOPTRunServiceBody113.Description
	variable.AOPTRunServiceBody113.ServiceID		= config.AOPTRunServiceBody113.ServiceID
	variable.AOPTRunServiceBody113.Token			= config.AOPTRunServiceBody113.Token
	variable.AOPTRunServiceBody113.StringInterface	= config.AOPTRunServiceBody113.StringInterface
}