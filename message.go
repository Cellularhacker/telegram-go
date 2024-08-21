package telegram

import (
	"fmt"
	"github.com/Cellularhacker/core-go"
	"github.com/Cellularhacker/logger"
	tb "gopkg.in/telebot.v3"
	"time"
)

func getToken() string {
	return accessToken
}

func SendMessage(message string, silent bool, chat ...tb.Recipient) {
	SendMessageAt(message, getNow(), silent, chat...)
}

func SendMessageAt(message string, at time.Time, silent bool, chat ...tb.Recipient) {
	if !core.IsProductionMode() || !enabled {
		return
	}
	var toChat tb.Recipient
	toChat = GetNormal()

	if len(chat) > 0 {
		toChat = chat[0]
	}

	msg := fmt.Sprintf("<%s> %s\n%s", serverAndNodeName, message, at.Format(time.RFC3339))
	logger.L.Debug("Sending telegram Message...")
	opt := &tb.SendOptions{
		DisableNotification: silent,
	}
	_, err := bot.Send(toChat, msg, opt)
	if err != nil {
		logger.L.Errorw(err.Error(), "func", "SendMessageAt()", "extra", "bot.Send(to, msg)", "to", toChat, "msg", msg)
		return
	}
	logger.L.Debug("[Telegram] message sent.")
}

func SendStarted(hostname string, localIP string, pubIP string) {
	msg := fmt.Sprintf("Server started successfully\nHostname:%s\nLocal IP:%s\nPublic IP:%s\n", hostname, localIP, pubIP)
	SendMessage(msg, false, GetMonitor())
}

func SendStopped(hostname string, localIP, pubIP string) {
	msg := fmt.Sprintf("Server stopping normally\nHostname:%s\nLocal IP:%s\nPublic IP:%s", hostname, localIP, pubIP)
	SendMessage(msg, false, GetMonitor())
}

func SendFailed(location string, err error) {
	msg := fmt.Sprintf("[ERROR/%s]\n=> %s", location, err)
	SendMessage(msg, false, GetMonitor())
}

func getNow() time.Time {
	return time.Now().In(core.GetLoc())
}
