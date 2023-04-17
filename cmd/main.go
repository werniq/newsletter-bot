package main

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"news-bot/Driver"
	"news-bot/app"
	"news-bot/internal/models"
	"news-bot/logger"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.NewLogger().Error("loading .env file", err)
		return
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		logger.NewLogger().Error("creating new telegram bot session ", err)
		return
	}

	bot.Debug = true

	db, err := Driver.OpenDb()
	if err != nil {
		logger.NewLogger().Error("opening database", err)
		return
	}

	application := app.NewApplication(bot, &models.DatabaseModel{DB: db})

	r := gin.Default()

	updatesChannel := bot.ListenForWebhook("http://localhost:8080/bot" + bot.Token)

	for update := range updatesChannel {
		application.StoreInfoInDB(update)
		application.Config(r)
	}

	err = r.Run(":8080")
	if err != nil {
		logger.NewLogger().Error("sending message", err)
		return
	}
}
