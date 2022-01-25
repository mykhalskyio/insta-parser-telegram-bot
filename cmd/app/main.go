package main

import (
	"log"
	"time"

	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/config"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/db"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/parser"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/telegram"
)

func main() {
	cfg := config.GetConfig()

	bot, err := telegram.NewBot(cfg.Telegram.Token)
	if err != nil {
		log.Fatalln("Telegram bot create error:", err)
	}

	db, err := db.NewConnect(cfg)
	if err != nil {
		log.Fatalln("BD error:", err)
	}

	err = db.MigrationInit()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		err := parser.Start(bot, db, cfg)
		if err != nil {
			log.Println("Parser error:", err)
		}
		time.Sleep(time.Minute * time.Duration(cfg.Parser.Minutes))
	}
}
