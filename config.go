package telegram

import (
	"fmt"
	"github.com/Cellularhacker/logger"
	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v3"
	"os"
	"time"
)

var (
	chatID        = ""
	monitorChatID = ""
	accessToken   = ""
)

var to *Chat
var toMonitor *Chat
var enabled = false

type Chat struct {
	id string
}

func (c Chat) Recipient() string {
	return c.id
}
func (c Chat) SetID(id string) {
	c.id = id
}
func (c Chat) String() string {
	return fmt.Sprintf("%s", c.id)
}

func init() {
	// MARK: <START>Initializing Variables
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	} else {
		logger.Info("Loaded environments from env file")
	}
	// MARK: <EMD>Initializing Variables

	// MARK: Applying environments
	chatID = os.Getenv(keyEnvChatID)
	monitorChatID = os.Getenv(keyEnvMonitorChatID)
	accessToken = os.Getenv(keyEnvAccessToken)

	logger.L.Info("Initializing telegram bot..")
	if chatID == "" {
		logger.L.Fatalf("'%s' missing.", keyEnvChatID)
	} else {
		to = &Chat{}
		to.SetID(chatID)
	}

	if monitorChatID == "" {
		// MARK: if monitorChatID is not specified, it will send as same as chatID
		logger.L.Warnf("'%s' missing. Default admin messages also send to the normal chat.", keyEnvMonitorChatID)
		toMonitor = &Chat{}
		toMonitor.SetID(chatID)
	} else {
		toMonitor = &Chat{}
		toMonitor.SetID(monitorChatID)
	}

	if accessToken == "" {
		logger.L.Fatalf("%s missing.", keyEnvAccessToken)
	}

	bot, err = tb.NewBot(tb.Settings{
		Token:  getToken(),
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		logger.L.Fatal(err.Error(), "func", "Init()", "extra", "tb.NewBot()")
	}

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
