package app

import (
	"duolingo-bot/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (app *Application) Help(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, I'm NewsLetter bot. I can send you news from different categories. To start using me, type /news.")
	if _, err := bot.Send(msg); err != nil {
		logger.NewLogger().Error("sending message ", err)
		return
	}
}
