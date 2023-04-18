package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"news-bot/logger"
	"strings"
)

// AddCategories function is used to add categories to the daily mailing.
func (app *Application) AddCategories(bot *tgbotapi.BotAPI, upd *tgbotapi.Update) {
	var err error
	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "What categories do you want to add to your daily mailing?")
	if _, err = bot.Send(msg); err != nil {
		logger.NewLogger().Error("sending message ", err)
		return
	}

	updateConfig := app.Bot.ListenForWebhook("http://localhost:8080/bot" + app.Bot.Token)

	update := <-updateConfig
	if update.Message == nil {
		return
	}

	categorys := "business" + " entertainment" + " general" + " health" + " science " + "sports" + " technology"

	args := strings.Split(update.Message.Text, " ")
	for i := 0; i < len(args); i++ {
		if strings.Contains(categorys, args[i]) {
			fmt.Println("ok")
		} else {
			msg.Text = "Wrong category name. Please choose from." + categorys
			if _, err = bot.Send(msg); err != nil {
				logger.NewLogger().Error("sending message ", err)
			}
			return
		}
	}

	err = app.Database.UpdateExistingDailyMailingCategories(update.Message.Chat.ID, args)
	if err != nil {
		logger.NewLogger().Error("storing daily mailing categories", err)
		return
	}

	fmt.Println("AddCategories command was used")
}
