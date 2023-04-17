package app

import (
	"duolingo-bot/internal/models"
	"duolingo-bot/logger"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
)

var uri = "https://newsapi.org/v2/everything"

// ReceiveNews function is used to receive news from the newsapi.org.
func (app *Application) ReceiveNews(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (*[]models.NewsApiResponse, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		logger.NewLogger().Error("creating new request", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	news := &[]models.NewsApiResponse{}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.NewLogger().Error("sending request", err)
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(&news)
	if err != nil {
		logger.NewLogger().Error("decoding response body", err)
		return nil, err
	}

	return news, nil
}
