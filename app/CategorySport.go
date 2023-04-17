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

// SportCategory function is used to send news from sport category
func (app *Application) SportCategory(bot *tgbotapi.BotAPI, upd *tgbotapi.Update) {
	rUri := uri + "?q=sports"
	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Here are the latest news from sport category: \n")

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

	msg.Text = msg.Text + news[0].Articles[0].Title + "\n" + news[0].Articles[0].Description + "\n" + news[0].Articles[0].URL + "\n\n"
	if _, err = bot.Send(msg); err != nil {
		logger.NewLogger().Error("sending message ", err)
		return
	}
	fmt.Println("Sport category function called")
}
