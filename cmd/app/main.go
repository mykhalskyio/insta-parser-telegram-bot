package main

import (
	"log"
	"time"

	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/config"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/db"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/instagram"
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
		errString := err.Error()
		bot.SendError(cfg.Telegram.User, "BD error:"+errString)
	}

	err = db.MigrationInit()
	if err != nil {
		errString := err.Error()
		bot.SendError(cfg.Telegram.User, "Migration error:"+errString)
	}

	insta := instagram.NewUser(cfg.Instagram.User, cfg.Instagram.Pass)

	for {
		currecntTime := getCurrentTime()
		hour := currecntTime.Hour()
		if (10 <= hour && hour <= 16) || (hour == 20) {
			err := parser.Start(insta, bot, db, cfg)
			if err != nil {
				errString := err.Error()
				bot.SendError(cfg.Telegram.User, "Parser error:"+errString)
			}
		}
		time.Sleep(time.Minute * time.Duration(cfg.Parser.Minutes))
	}
}

func getCurrentTime() time.Time {
	now := time.Now()
	loc, _ := time.LoadLocation("Europe/Kiev")
	return now.In(loc)
}
