package slogger

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

const (
	_handlerTypeText = 0
	_handlerTypeJSON = 1
)

type Logger struct {
	lg *slog.Logger
}

func NewLogger(level, handlerType int) *Logger {
	opts := &slog.HandlerOptions{
		Level: slogLeveler(level),
	}

	return &Logger{
		lg: slog.New(slogHandler(handlerType, opts)),
	}
}

func slogLeveler(level int) slog.Leveler {
	switch level {
	case int(slog.LevelDebug):
		return slog.LevelDebug
	case int(slog.LevelInfo):
		return slog.LevelInfo
	case int(slog.LevelWarn):
		return slog.LevelWarn
	case int(slog.LevelError):
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func slogHandler(handlerType int, opts *slog.HandlerOptions) slog.Handler {
	switch handlerType {
	case _handlerTypeText:
		return slog.NewTextHandler(os.Stdout, opts)
	case _handlerTypeJSON:
		return slog.NewJSONHandler(os.Stdout, opts)
	default:
		return slog.NewTextHandler(os.Stdout, opts)
	}
}

func (r *Logger) Debug(message interface{}, args ...interface{}) {
	if len(args) == 0 {
		r.lg.Debug(r.toString(message))
		return
	}

	out := []interface{}{"args"}
	r.lg.Debug(r.toString(message), append(out, args...)...)
}

func (r *Logger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		r.lg.Debug(r.toString(message))
		return
	}

	out := []interface{}{"args"}
	r.lg.Info(message, append(out, args...)...)
}

func (r *Logger) Warn(message string, args ...interface{}) {
	if len(args) == 0 {
		r.lg.Debug(r.toString(message))
		return
	}

	out := []interface{}{"args"}
	r.lg.Warn(message, append(out, args...)...)
}

func (r *Logger) Error(message interface{}, args ...interface{}) {
	if len(args) == 0 {
		r.lg.Debug(r.toString(message))
		return
	}

	out := []interface{}{"args"}
	r.lg.Error(r.toString(message), append(out, args...)...)
}

func (r *Logger) Fatal(message interface{}, args ...interface{}) {
	if len(args) == 0 {
		r.lg.Debug(r.toString(message))
		return
	}

	out := []interface{}{"args"}
	r.lg.Error(r.toString(message), append(out, args...)...)

	os.Exit(1)
}

func (r *Logger) toString(message interface{}) string {
	var msg string

	switch v := message.(type) {
	case error:
		msg = v.Error()
	case string:
		msg = v
	default:
		msg = fmt.Sprintf("something: %s", v)
	}

	return msg
}

func (r *Logger) Timing(message string, starting time.Time) {
	r.lg.Debug(message, "time", time.Now().Sub(starting).String())
}
