package logger

import (
	"log"
	"os"
	"time"
)

// Logger struct to encapsulate log handling
type Logger struct {
	file *os.File
}

// NewLogger initializes a new Logger
func NewLogger(logFile string) (*Logger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)
	return &Logger{file: file}, nil
}

// LogError logs an error message with a timestamp
func (l *Logger) LogError(err error) {
	if err != nil {
		log.Printf("%s - ERROR: %s\n", time.Now().Format(time.RFC3339), err.Error())
	}
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}
