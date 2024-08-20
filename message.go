package telegram

import (
	"fmt"
	"github.com/Cellularhacker/core-go"
	"github.com/Cellularhacker/logger"
	"time"
)

const (
	keyEnvChatID        = "TELEGRAM_CHAT_ID"
	keyEnvMonitorChatID = "TELEGRAM_MONITOR_CHAT_ID"
	keyEnvAccessToken   = "TELEGRAM_ACCESS_TOKEN"
)

func getToken() string {
	return accessToken
}

func SendMessage(message string, chat ...*Chat) {
	SendMessageAt(message, getNow(), chat...)
}

func SendMessageAt(message string, at time.Time, chat ...*Chat) {
	if !core.IsProductionMode() || !enabled {
		return
	}
	toChat := to
	if len(chat) > 0 {
		toChat = chat[0]
	}

	msg := fmt.Sprintf("<%s> %s\n%s", serverAndNodeName, message, at.Format(time.RFC3339))
	logger.L.Debug("Sending telegram Message...")
	_, err := bot.Send(*toChat, msg)
	if err != nil {
		logger.L.Errorw(err.Error(), "func", "SendMessageAt()", "extra", "bot.Send(to, msg)", "to", toChat, "msg", msg)
		return
	}
	logger.L.Debug("[Telegram] message sent.")
}

func SendStarted(hostname string, localIP string, pubIP string) {
	msg := fmt.Sprintf("Server started successfully\nHostname:%s\nLocal IP:%s\nPublic IP:%s\n", hostname, localIP, pubIP)
	SendMessage(msg, toMonitor)
}

func SendStopped(hostname string, localIP, pubIP string) {
	msg := fmt.Sprintf("Server stopping normally\nHostname:%s\nLocal IP:%s\nPublic IP:%s", hostname, localIP, pubIP)
	SendMessage(msg, toMonitor)
}

func SendFailed(location string, err error) {
	msg := fmt.Sprintf("[ERROR/%s]\n=> %s", location, err)
	SendMessage(msg, toMonitor)
}

func getNow() time.Time {
	return time.Now().In(core.GetLoc())
}
