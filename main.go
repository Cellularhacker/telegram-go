package telegram

import (
	"fmt"
	"github.com/Cellularhacker/logger"
	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v3"
	"os"
	"time"
)

const (
	keyEnvChatID        = "TELEGRAM_CHAT_ID"
	keyEnvMonitorChatID = "TELEGRAM_MONITOR_CHAT_ID"
	keyEnvAccessToken   = "TELEGRAM_ACCESS_TOKEN"
	keyModeProduction   = "Production"
	keyModeDevelopment  = "Development"
)

var to *Chat
var enabled = false

var (
	chatID        = ""
	monitorChatID = ""
	accessToken   = ""
)

type Chat struct{}
type MonitorChat struct{}

func (Chat) Recipient() string {
	return chatID
}
func (MonitorChat) Recipient() string {
	return monitorChatID
}

func getToken() string {
	return accessToken
}

var bot *tb.Bot

func init() {
	// MARK: <START>Initializing Variables
	envFileLoaded := false
	for i := range os.Args {
		if os.Args[i] == "env" {
			fileName := os.Args[i+1]
			err := godotenv.Load(fileName)
			if err != nil {
				logger.Fatal("Error loading .env file", err)
			}
			envFileLoaded = true
			break
		}
	}

	if !envFileLoaded {
		err := godotenv.Load()
		if err != nil {
			logger.Fatal("Error loading .env file")
		} else {
			logger.Info("Loaded environments from env file")
		}
	}
	// MARK: <EMD>Initializing Variables

	// MARK: Applying environments
	chatID = os.Getenv(keyEnvChatID)
	monitorChatID = os.Getenv(keyEnvMonitorChatID)
	accessToken = os.Getenv(keyEnvAccessToken)

	logger.L.Info("Initializing telegram bot..")
	if chatID == "" {
		logger.L.Fatalf("'%s' missing.", keyEnvChatID)
	}
	if monitorChatID == "" {
		// MARK: if monitorChatID is not specified, it will send as same as chatID
		logger.L.Warnf("'%s' missing.", keyEnvMonitorChatID)
		chatID = monitorChatID
	}
	if accessToken == "" {
		logger.L.Fatalf("%s missing.", keyEnvAccessToken)
	}

	var err error

	bot, err = tb.NewBot(tb.Settings{
		Token:  getToken(),
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		logger.L.Fatal(err.Error(), "func", "Init()", "extra", "tb.NewBot()")
	}

	to = &Chat{}
	enabled = true
}

func SendMessage(message string) {
	SendMessageAt(message, getNow())
}

func SendMessageAt(message string, at time.Time) {
	if !IsProductionMode() || !enabled {
		return
	}
	msg := fmt.Sprintf("<ch-api> %s\n%s", message, at.Format(time.RFC822))
	logger.L.Debug("Sending telegram Message...")
	_, err := bot.Send(to, msg)
	if err != nil {
		logger.L.Errorw(err.Error(), "func", "SendMessageAt()", "extra", "bot.Send(to, msg)", "to", to.Recipient(), "msg", msg)
		return
	}
	logger.L.Debug("[Telegram] message sent.")
}

func SendStarted(hostname string, localIP string, pubIP string) {
	msg := fmt.Sprintf("Server started successfully\nHostname:%s\nLocal IP:%s\nPublic IP:%s\n", hostname, localIP, pubIP)
	SendMessage(msg)
}

func SendStopped(hostname string, localIP, pubIP string) {
	msg := fmt.Sprintf("Server stopping normally\nHostname:%s\nLocal IP:%s\nPublic IP:%s", hostname, localIP, pubIP)
	SendMessage(msg)
}

func SendFailed(location string, err error) {
	msg := fmt.Sprintf("[ERROR/%s]\n=> %s", location, err)
	SendMessage(msg)
}

func getNow() time.Time {
	return time.Now().In(config.Loc)
}
