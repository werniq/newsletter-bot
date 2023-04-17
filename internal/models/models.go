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
