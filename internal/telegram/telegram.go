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

// send media to channel
func (bot *TelegramBot) SendToChannel(channelName string, data tgbotapi.RequestFileData, dataType int) error {
	if dataType == 1 {
		_, err := bot.Api.Send(tgbotapi.NewPhotoToChannel("@"+channelName, data))
		if err != nil {
			return err
		}
	} else {
		videoConfig := tgbotapi.NewVideo(0, data)
		videoConfig.ChannelUsername = "@" + channelName
		_, err := bot.Api.Send(videoConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bot *TelegramBot) SendError(user int64, err string) {
	msg := tgbotapi.NewMessage(user, err)
	bot.Api.Send(msg)
}
