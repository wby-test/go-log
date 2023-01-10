package main

import (
	"errors"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			a.Key = "time"
			a.Value = slog.StringValue("test")
			return a
		},
	}
	textHandler := opts.NewTextHandler(os.Stdout)
	logger := slog.New(textHandler)
	// TODO: traceID
	//ctx := slog.NewContext(context.Background(), logger)
	//logger = slog.FromContext(ctx)

	logger.Info("Go is the best language!")
	logger.Error("Go is the worst language!", errors.New("trash"))
	logger.Info("x", slog.String("tracID", "xxxx"))
}
