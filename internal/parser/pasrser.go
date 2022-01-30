package parser

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/config"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/db"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/instagram"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/telegram"
)

// parse instgram storis
func Start(insta *instagram.InstaUser, bot *telegram.TelegramBot, dbpg *db.Postgres, cfg *config.Config) error {

	err := insta.User.Login()
	if err != nil {
		fmt.Println("Loggin error:", err)
		return err
	}
	defer insta.User.Logout()

	stories, err := insta.GetUserStories(cfg.Instagram.UserParse)
	if err != nil {
		return err
	}
	for _, storis := range stories {
		storisId := storis.GetID()

		result := dbpg.Check(storisId)
		if result {
			continue
		}

		media, _ := storis.Download()

		photoBytes := tgbotapi.FileBytes{
			Name:  "media",
			Bytes: media,
		}
		bot.SendToChannel(cfg.Telegram.Channel, photoBytes, storis.MediaType)
	}
	return nil
}
