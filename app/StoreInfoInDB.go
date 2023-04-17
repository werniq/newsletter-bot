package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"news-bot/logger"
)

// StoreInfoInDB function is used to store info about each message in database.
func (app *Application) StoreInfoInDB(update tgbotapi.Update) {
	err := app.Database.StoreInfoInDatabase(update)
	if err != nil {
		logger.NewLogger().Println("storing info in database", err)
		return
	}
	fmt.Println("message saved")
}
