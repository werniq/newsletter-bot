package models

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

type DatabaseModel struct {
	DB *sql.DB
}

// DailyMailingCategories is a struct which later I will use for user shortcut search preferences
// for example: user wants to search news by keyword "football" and "sport", so in morning he will receive news about football and sport.
type DailyMailingCategories struct {
	ChatID         int64
	NewsCategories []string
}

type Message struct {
	ID        int       `json:"id"`
	ChatID    int64     `json:"chat_id"`
	Content   string    `json:"content"`
	AuthorID  int64     `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

// StoreInfoInDatabase function is used to store the user's information in the database.
func (m *DatabaseModel) StoreInfoInDatabase(upd tgbotapi.Update) error {
	mes := &Message{
		ChatID:    upd.Message.Chat.ID,
		Content:   upd.Message.Text,
		AuthorID:  upd.Message.From.ID,
		CreatedAt: time.Now(),
	}

	stmt := `
			INSERT INTO 
			    messages 
			    (chat_id, content, author_id, created_at) 
			VALUES 
			    ($1, $2, $3, $4)`

	_, err := m.DB.Exec(stmt, mes.ChatID, mes.Content, mes.AuthorID, mes.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetDailyMailingCategories function is used to get user's preferences for daily mailing.
func (m *DatabaseModel) GetDailyMailingCategories(chatID int64) ([]string, error) {
	stmt := `
		SELECT 
		    news_categories
		FROM 
		    daily_mailing_categories 
		WHERE 
		    chat_id = $1`
	var newsCategories []string
	err := m.DB.QueryRow(stmt, chatID).Scan(&newsCategories)
	if err != nil {
		return nil, err
	}

	return newsCategories, nil
}

// StoreDailyMailingCategories function is used to store user's preferences for daily mailing.
func (m *DatabaseModel) StoreDailyMailingCategories(chatID int64, newsCategories []string) error {
	stmt := `
		INSERT INTO 
		    daily_mailing_categories 
		    (chat_id, news_categories) 
		VALUES 
		    ($1, $2)`

	_, err := m.DB.Exec(stmt, chatID, newsCategories)
	if err != nil {
		return err
	}

	return nil
}

// UpdateExistingDailyMailingCategories function is used to update user's preferences for daily mailing.
func (m *DatabaseModel) UpdateExistingDailyMailingCategories(chatID int64, newsCategories []string) error {
	categories, err := m.GetDailyMailingCategories(chatID)
	if err != nil {
		return err
	}
	for _, v := range newsCategories {
		categories = append(categories, v)
	}

	stmt := `
		UPDATE 
		    daily_mailing_categories 
		SET 
		    news_categories = $1 
		WHERE 
		    chat_id = $2`

	_, err = m.DB.Exec(stmt, categories, chatID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllMailingSubscrpiptedUsers function is used to get all users who subscribed to daily mailing.
func (m *DatabaseModel) GetAllMailingSubscrpiptedUsers() ([]int64, error) {
	stmt := `
		SELECT 
		    chat_id
		FROM 
		    daily_mailing_categories`
	var chatIDs []int64
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var chatID int64
		err := rows.Scan(&chatID)
		if err != nil {
			return nil, err
		}
		chatIDs = append(chatIDs, chatID)
	}

	return chatIDs, nil
}
