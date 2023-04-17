package logger

import (
	"log"
	"os"
)

// Logger is a wrapper for the log package
type Logger struct {
	*log.Logger
}

// NewLogger creates a new logger
func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "ERROR: \t", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile)}
}

// Error logs the error and the reason
func (l *Logger) Error(reason string, err error) {
	l.Printf("Error: %s, While: %s\n", err, reason)
}
