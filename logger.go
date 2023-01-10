package main

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

var logger *slog.Logger

func GetLogger(level slog.Level) *slog.Logger {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			a.Key = "time"
			a.Value = slog.StringValue("test")
			return a
		},
	}

	textHandler := opts.NewTextHandler(os.Stdout)
	logger = slog.New(textHandler)

	return logger
}

func Info(ctx context.Context, msg string, args ...any) {
	logger.WithContext(ctx).Info(msg, args)
}
