package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// bot struct
type TelegramBot struct {
	Api *tgbotapi.BotAPI
}

// create bot-api
func NewBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		Api: bot,
	}, nil
}
