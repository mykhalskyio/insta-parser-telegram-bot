package main

import (
	"log"

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

	log.Println(db.CheckTable())

	log.Println("Parser error:", parser.Parse(bot, db, cfg))

}
