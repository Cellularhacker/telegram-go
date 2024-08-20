package telegram

import (
	"github.com/Cellularhacker/logger"
	tb "gopkg.in/telebot.v3"
	"time"
)

var bot *tb.Bot

var (
	chatID            = ""
	monitorChatID     = ""
	accessToken       = ""
	serverAndNodeName = ""
)

func Init(ServerAndNodeName string, AccessToken string, ChatID string, MonitorChatID ...string) {
	// MARK: Applying environments

	logger.L.Info("Initializing telegram bot..")
	serverAndNodeName = ServerAndNodeName
	if len(ServerAndNodeName) <= 0 {
		logger.L.Fatal("ServerAndNodeName is empty")
	}

	if ChatID == "" {
		logger.L.Fatalf("'ChatID' missing.")
	} else {
		to = &Chat{}
		to.SetID(ChatID)
		chatID = ChatID
	}

	if len(MonitorChatID) <= 0 {
		// MARK: if monitorChatID is not specified, it will send as same as chatID
		logger.L.Warnf("'MonitorChatID' missing. Default admin messages also send to the normal chat.")
		toMonitor = &Chat{}
		monitorChatID = chatID
		toMonitor.SetID(monitorChatID)
	} else {
		toMonitor = &Chat{}
		toMonitor.SetID(MonitorChatID[0])
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
		logger.L.Fatal(err.Error(), "func", "Init()", "extra", "tb.NewBot()")
	}

	bot = tbBot
	enabled = true
}

//
// MARK: Utils

func GetNormal() *Chat {
	return to
}
func GetMonitor() *Chat {
	return toMonitor
}
