package parser

import (
	"fmt"

	"github.com/Davincible/goinsta"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/config"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/db"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/telegram"
)

// parse instgram storis
func Start(bot *telegram.TelegramBot, dbpg *db.Postgres, cfg *config.Config) error {
	user := goinsta.New(cfg.Instagram.User, cfg.Instagram.Pass)

	err := user.Login()
	if err != nil {
		fmt.Println("Loggin error:", err)
		return err
	}
	defer user.Logout()

	profile, _ := user.VisitProfile(cfg.Telegram.UserParse)

	stories := profile.Stories.Reel
	for _, storis := range stories.Items {
		storisId := storis.GetID()

		result := dbpg.Check(storisId)
		if result {
			continue
		}

		photo, _ := storis.Download()

		photoBytes := tgbotapi.FileBytes{
			Name:  "picture",
			Bytes: photo,
		}
		bot.SendToChannel(cfg.Telegram.Channel, photoBytes)
	}
	return nil
}
