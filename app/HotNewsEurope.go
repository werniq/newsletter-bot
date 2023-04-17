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

// HotNewsEu function is used to send the hottest news from Europe
func (app *Application) HotNewsEu(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "The hottest news right now: \n")
	var err error
	req, err := http.NewRequest("GET", "https://newsapi.org/v2/top-headlines?country=eu", nil)
	if err != nil {
		logger.NewLogger().Error("creating new request", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Key", os.Getenv("NEWS_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.NewLogger().Error("sending request", err)
		return
	}

	var news []models.NewsApiResponse
	err = json.NewDecoder(res.Body).Decode(&news)
	if err != nil {
		logger.NewLogger().Error("decoding response body", err)
		return
	}
	for i := 0; i < len(news); i++ {
		msg.Text += news[i].Articles[i].Title + "\n" + news[i].Articles[i].Description + "\n" + news[i].Articles[i].URL + "\n\n"
	}

	if _, err = bot.Send(msg); err != nil {
		logger.NewLogger().Error("sending message ", err)
		return
	}
	fmt.Println("Hot news function called")
}
