package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"news-bot/logger"
)

// Categories function is used to list all available categories
func (app *Application) Categories(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available categories: \n")
	var err error
	for i := 0; i < len(categories); i++ {
		msg.Text += categories[i] + ", " + "\n"
	}
	if _, err = bot.Send(msg); err != nil {
		logger.NewLogger().Error("sending message ", err)
		return
	}
}
