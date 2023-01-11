package logger

import (
	"errors"
	"testing"

	"golang.org/x/exp/slog"
)

func TestStand_Debug(t *testing.T) {
	var logger Logger
	logger = NewStandLogger(slog.LevelDebug)
	logger.Error(errors.New("test"), "info", "xxxx")
	logger.Debug("debug", "xxxx")
	logger.Warn("warn", "xxxx")
}
