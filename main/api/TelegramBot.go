package api

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	"net/url"
	"thanhgit.com/mywebhook/main/utils"
	"thanhgit.com/mywebhook/main/variable"
	"time"
)

type TelegramBot struct {
	instance *tgbotapi.BotAPI
}

func (bot *TelegramBot) GetInstance() *tgbotapi.BotAPI  {
	if bot.instance == nil {
		proxyUrl, err := url.Parse(variable.GetInstance().Proxy)
		client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
		client.Timeout = 10 * time.Second
		var _bot *tgbotapi.BotAPI
		_bot, err = tgbotapi.NewBotAPIWithClient(variable.GetInstance().TelegramToken, client)
		if err != nil {
			//panic(err) // You should add better error handling than this!
			println(err.Error())
			_bot, err = tgbotapi.NewBotAPI(variable.GetInstance().TelegramToken)
		}

		//_bot.Debug = true // Has the library display every request and response.
		bot.instance = _bot

		utils.PrintLog("Authorized on account " +_bot.Self.UserName)
	}

	return bot.instance
}

func (bot *TelegramBot) GetInstanceWithToken(_token string) *tgbotapi.BotAPI  {
	if bot.instance == nil {
		proxyUrl, err := url.Parse(variable.GetInstance().Proxy)
		client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
		client.Timeout = 10 * time.Second
		var _bot *tgbotapi.BotAPI
		_bot, err = tgbotapi.NewBotAPIWithClient(_token, client)
		if err != nil {
			//panic(err) // You should add better error handling than this!
			println(err.Error())
			_bot, err = tgbotapi.NewBotAPI(_token)
		}

		//_bot.Debug = true // Has the library display every request and response.
		bot.instance = _bot

		utils.PrintLog("Authorized on account " + _bot.Self.UserName)
	}

	return bot.instance
}

func (bot *TelegramBot) SendMessage(chatRomdId int64, msg string) {
	hiMsg := tgbotapi.NewMessage(chatRomdId, msg)
	_, err := bot.GetInstance().Send(hiMsg)

	if err != nil {
		println(err)
	}
}

// Send Message to kibana alert channel
func (bot *TelegramBot) SendMessageToDefaultChannel(msg string)  {
	chatID := variable.GetInstance().ChatID
		bot.SendMessage(chatID, msg)
}

func (bot *TelegramBot) BotAutomaticReply() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetInstance().GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		fmt.Println("Chat", update.ChannelPost)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		msg.Text = "Hello everyone!"

		if _, err := bot.GetInstance().Send(msg); err != nil {
			panic(err)
		}
	}
}