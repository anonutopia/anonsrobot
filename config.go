package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct holds all our configuration
type Config struct {
	Dev            bool   `json:"dev"`
	Debug          bool   `json:"debug"`
	TelegramAPIKey string `json:"telegram_api_key"`
	PostgreSQL     string `json:"postgre_sql"`
}

// Load method loads configuration file to Config struct
func (c *Config) load(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		log.Println(err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		log.Println(err)
	}
}

func initConfig() *Config {
	c := &Config{}
	c.load("config.json")
	return c
}
