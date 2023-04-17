package app

import (
	"duolingo-bot/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

var (
	categories = []string{"business", "entertainment", "general", "health", "science", "sports", "technology"}
)

// Categories function is used to clarify that the user has entered the correct categories.
func (app *Application) Categories(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	var err error
	if update.Message.CommandArguments() != "" {
		arg := update.Message.CommandArguments()
		args := strings.Split(arg, " ")
		if len(args) != 1 {
			msg.Text = "You should provide categories 1 by 1."
			if _, err = bot.Send(msg); err != nil {
				logger.NewLogger().Error("sending message ", err)
				return
			}
		}
		for i := 0; i < len(categories); i++ {
			if strings.Contains(arg, categories[i]) == false {
				if _, err = bot.Send(msg); err != nil {
					logger.NewLogger().Error("sending message ", err)
					return
				}
			}
		}
	} else {
		msg.Text = "You should provide at least 1 category."
		if _, err := bot.Send(msg); err != nil {
			logger.NewLogger().Error("sending message ", err)
			return
		}
	}
}
