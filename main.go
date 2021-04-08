package main

import (
	"log"

	"gopkg.in/tucnak/telebot.v2"
)

var conf *Config

var bot *telebot.Bot

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	bot = initTelegramBot()

	logTelegram("AnonsRobot daemon successfully started. 🚀")
	log.Println("AnonsRobot daemon successfully started. 🚀")

	bot.Start()
}
