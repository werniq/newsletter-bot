package Driver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"news-bot/logger"
	"os"
)

// OpenDb function is used to open the database connection.
func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_DSN"))
	if err != nil {
		logger.NewLogger().Error("opening database connection ", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.NewLogger().Error("pinging database ", err)
		return nil, err
	}
	return db, nil
}

// CloseDb function is used to close the database connection.
func CloseDb(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		logger.NewLogger().Error("closing database connection ", err)
		return err
	}
	return nil
}
