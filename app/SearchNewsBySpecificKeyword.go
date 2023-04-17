package app

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"news-bot/internal/models"
	"news-bot/logger"
	"os"
)

// SearchNewsBySpecificKeyword function is used to search news by specific keyword.
func (app *Application) SearchNewsBySpecificKeyword(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	var err error

	if update.Message.CommandArguments() == "" {
		msg.Text = "Please, enter a keyword"
		if _, err = bot.Send(msg); err != nil {
			logger.NewLogger().Error("sending message ", err)
			return
		}
		return
	}

	keyword := update.Message.CommandArguments()
	rUri := "https://newsapi.org/v2/everything?q=" + keyword
	req, err := http.NewRequest("GET", rUri, nil)
	if err != nil {
		logger.NewLogger().Error("creating request ", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Key", os.Getenv("NEWS_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.NewLogger().Error("sending request ", err)
		return
	}

	var news []models.NewsApiResponse

	err = json.NewDecoder(res.Body).Decode(&news)
	if err != nil {
		logger.NewLogger().Error("decoding response body ", err)
		return
	}
	msg.Text = fmt.Sprintf("The latest news about "+keyword+":\n\n", "are ", news[0].Articles[0].Title, " ", news[0].Articles[0].Description, " ", news[0].Articles[0].URL)
	_, err = bot.Send(msg)
	if err != nil {
		logger.NewLogger().Error("sending message", err)
		return
	}
	fmt.Println("Search news by specific keyword function called")
}
