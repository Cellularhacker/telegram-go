package telegram

import (
	"github.com/Cellularhacker/logger-go"
	tb "gopkg.in/telebot.v3"
	"time"
)

var toNormal NormalChat
var toMonitor MonitorChat

var bot *tb.Bot

var (
	chatID            = ""
	monitorChatID     = ""
	accessToken       = ""
	serverAndNodeName = ""
)

func Init(ServerAndNodeName string, AccessToken string, ChatID string, MonitorChatID ...string) {
	// MARK: Applying environments

	logger.Info("Initializing telegram bot..")
	serverAndNodeName = ServerAndNodeName
	if len(ServerAndNodeName) <= 0 {
		logger.L.Fatal("ServerAndNodeName is empty")
	}

	if ChatID == "" {
		logger.Fatalf("'ChatID' missing.")
	} else {
		chatID = ChatID
	}

	if len(MonitorChatID) <= 0 {
		// MARK: if monitorChatID is not specified, it will send as same as chatID
		logger.Warnf("'MonitorChatID' missing. Default admin messages also send to the normal chat.")
		monitorChatID = chatID
	} else {
		monitorChatID = MonitorChatID[0]
	}

	if AccessToken == "" {
		logger.L.Fatalf("'AccessToken' missing.")
	}
	accessToken = AccessToken

	tbBot, err := tb.NewBot(tb.Settings{
		Token:  getToken(),
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		logger.Fatal(err.Error(), "func", "Init()", "extra", "tb.NewBot()")
	}

	bot = tbBot
	enabled = true
}

//
// MARK: Utils

func GetNormal() tb.Recipient {
	return toNormal
}
func GetMonitor() tb.Recipient {
	return toMonitor
}
