package main

import (
	"log"

	"gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

var conf *Config

var bot *telebot.Bot

var db *gorm.DB

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	bot = initTelegramBot()

	db = initDb()

	logTelegram("AnonsRobot daemon successfully started. 🚀")
	log.Println("AnonsRobot daemon successfully started. 🚀")

	bot.Start()
}
