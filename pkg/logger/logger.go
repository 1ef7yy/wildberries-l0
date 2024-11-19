package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger() Logger {
	// логгер в формате json
	return Logger{
		log: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (l *Logger) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.log.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Error(msg)
}
