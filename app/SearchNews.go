package app

import (
	"duolingo-bot/internal/models"
	"duolingo-bot/logger"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"os"
	"strings"
)

func (app *Application) SearchNews(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, enter the categories you want to get news from. Available categories: business, entertainment, general, health, science, sports, technology.")
	if _, err := bot.Send(msg); err != nil {
		return
	}

	// listen for user's input
	// if user's input is correct, send news
	// if user's input is incorrect, send message about incorrect input
	updatesChannel := bot.ListenForWebhook("/" + bot.Token)

	// retrieve latest update

	for {
		upd := <-updatesChannel

		if upd.Message.Text != "" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "You should choose from available categories.")
			if _, err := bot.Send(msg); err != nil {
				logger.NewLogger().Error("sending message", err)
				return
			}
			args := strings.Split(update.Message.Text, " ")

			for i := 0; i < len(args); i++ {
				if args[i] == "business" || args[i] == "entertainment" || args[i] == "general" || args[i] == "health" || args[i] == "science" || args[i] == "sports" || args[i] == "technology" {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "You have chosen "+args[i]+" category.")
					if _, err := bot.Send(msg); err != nil {
						logger.NewLogger().Error("sending message", err)
						return
					}
					rUri := uri + "?q=" + args[i]

					// arguments are valid
					// now we should send news

					// request to news api
					req, err := http.NewRequest("GET", rUri, nil)
					if err != nil {
						logger.NewLogger().Error("request to news api", err)
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
					// and send news to user
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Here are the news:")
					_, err = bot.Send(msg)
					if err != nil {
						logger.NewLogger().Error("sending message", err)
						return
					}
					b := 0

					msg.Text += news[b].Articles[b].Title + "\n" + news[b].Articles[b].Description + "\n" + news[b].Articles[b].URL + "\n\n"
					_, err = bot.Send(msg)
					if err != nil {
						logger.NewLogger().Error("sending message", err)
						return
					}
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "You should choose from available categories, or type /exit to exit.")
					if _, err := bot.Send(msg); err != nil {
						logger.NewLogger().Error("sending message", err)
						return
					}
				}
			}
			break
		} else if upd.Message.Text == "/exit" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "You have exited the search.")
			if _, err := bot.Send(msg); err != nil {
				logger.NewLogger().Error("sending message", err)
				return
			}
			break
		}
		break
	}
	fmt.Println("SearchNews command was used")
}
