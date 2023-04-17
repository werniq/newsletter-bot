package app

import (
	"duolingo-bot/internal/models"
	"duolingo-bot/logger"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Application struct {
	Bot      *tgbotapi.BotAPI
	Database *models.DatabaseModel
}

// NewApplication is used to create new application.
func NewApplication(bot *tgbotapi.BotAPI, db *models.DatabaseModel) *Application {
	return &Application{
		Bot:      bot,
		Database: db,
	}
}

// Config is used to configure application handlers.
func (app *Application) Config(r *gin.Engine) {
	r.POST("/bot"+app.Bot.Token, app.Configure)
}

// Configure is used to configure application handlers.
func (app *Application) Configure(c *gin.Context) {
	// handle incoming messages
	update := &tgbotapi.Update{}
	err := c.ShouldBindJSON(&update)
	if err != nil {
		logger.NewLogger().Error("decoding request body ", err)
		return
	}

	// handle update
	if update.Message.Text == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "invalid command")
		_, err = app.Bot.Send(msg)
		if err != nil {
			logger.NewLogger().Error("sending message", err)
			return
		}
	}

	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "/start":
			msg.Text = "Hello, I'm NewsLetter bot. I can send you news from different categories. To start using me, type /news."
			_, err = app.Bot.Send(msg)
			if err != nil {
				logger.NewLogger().Error("sending message", err)
				return
			}
		case "/help":
			msg.Text = "Hello, I'm NewsLetter bot. You can use me by choosing one of the following commands: /news, /hotnews."
			_, err = app.Bot.Send(msg)
			if err != nil {
				logger.NewLogger().Error("sending message", err)
				return
			}
		case "/news":
			app.SearchNews(app.Bot, update)
		case "/randomnews":
		case "/hotnews":

		case "/sportnews":
		case "/politicsnews":
		case "/economicsnews":
		case "/scienceandtechnews":
		case "/healthnews":
		case "/entertainmentnews":
		case "/travelnews":
		case "/lifestylenews":

		case "/foodnews":

		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
			_, err = app.Bot.Send(msg)
			if err != nil {
				logger.NewLogger().Error("sending message ", err)
				return
			}
		}
	}
}
