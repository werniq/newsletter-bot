package app

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"news-bot/internal/models"
	"news-bot/logger"
	"os"
	"time"
)

// DailyMorningMailing function is used to send daily morning mailing to all subscribed users.
func (app *Application) DailyMorningMailing(bot *tgbotapi.BotAPI, upd tgbotapi.Update) {
	loc, err := time.LoadLocation("EUROPE/WARSAW")
	if err != nil {
		logger.NewLogger().Error("loading location", err)
		return
	}

	ticker := time.NewTicker(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, loc).Sub(time.Now()))

	subscriptedUsers, err := app.Database.GetAllMailingSubscrpiptedUsers()
	if err != nil {
		logger.NewLogger().Error("getting all mailing subscribed users", err)
		return
	}

	go func() {
		for {
			// wait for the next tick
			<-ticker.C

			// create new message

			for i := 0; i <= len(subscriptedUsers)-1; i++ {
				msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Good morning! Here is your daily news: \n")
				_, err = bot.Send(msg)
				if err != nil {
					logger.NewLogger().Error("sending message ", err)
					return
				}
				categories, err := app.Database.GetDailyMailingCategories(upd.Message.Chat.ID)
				if err != nil {
					logger.NewLogger().Error("getting categories for daily mailing", err)
					return
				}

				for i := 0; i < len(categories); i++ {
					category := categories[i]

					req, err := http.NewRequest("GET", uri+"?q="+category, nil)
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
						_, err = bot.Send(msg)
						if err != nil {
							logger.NewLogger().Error("sending message ", err)
							return
						}
					}
				}

				fmt.Println("Daily morning mailing function called")

			}
		}
	}()
	select {}
}
