package main

import (
	"strings"

	"github.com/bykovme/gotrans"
	"gopkg.in/tucnak/telebot.v2"
)

func initCommands() {
	bot.Handle(telebot.OnUserJoined, userJoined)
	bot.Handle(telebot.OnText, unknownCommand)
}

func userJoined(m *telebot.Message) {
	bot.Send(m.Sender, gotrans.T("welcome"))
}

func unknownCommand(m *telebot.Message) {
	if m.Private() {
		bot.Send(m.Sender, gotrans.T("unknown"))
	} else if strings.HasPrefix(m.Text, "/") {
		bot.Send(m.Chat, gotrans.T("unknown"))
	}
}
