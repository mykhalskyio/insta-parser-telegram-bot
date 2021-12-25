package parser

import (
	"log"
	"time"

	"github.com/Davincible/goinsta"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/config"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/db"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/telegram"
)

// parse instgram storis (не дуже вигляжає, але працює)
func Parse(bot *telegram.TelegramBot, dbpg *db.Postgres, cfg *config.Config) error {
	user := goinsta.New(cfg.Instagram.User, cfg.Instagram.Pass)

	err := user.Login()
	if err != nil {
		return err
	}

	profile, _ := user.VisitProfile(cfg.Instagram.UserParse)

	for {
		log.Println("Start cicle")

		log.Println("Get stories")
		stories := profile.Stories.Reel
		for _, storis := range stories.Items {
			storisId := storis.GetID()

			result := dbpg.Check(storisId)
			if result {
				continue
			}

			photo, _ := storis.Download()

			log.Println("Sent storis")
			bot.Api.Send(tgbotapi.NewPhotoToChannel("@"+cfg.Instagram.Channel, tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: photo,
			}))
			log.Println("Sent - Ok")
		}
		log.Println("Cicle - ok")
		time.Sleep(time.Minute * time.Duration(cfg.Instagram.Minutes))
	}
}
