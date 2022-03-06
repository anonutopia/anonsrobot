package main

import (
	"github.com/bykovme/gotrans"
	"gopkg.in/tucnak/telebot.v2"
)

func initLangs() {
	gotrans.InitLocales("langs")
	gotrans.SetDefaultLocale("en")
}

func tr(text string, lang string) string {
	return gotrans.Tr(lang, text)
}

func trG(text string, m *telebot.Message) string {
	var lang string

	if m.Chat.ID == TelKriptokuna {
		lang = "hr"
	} else {
		lang = "en"
	}

	return tr(text, lang)
}
