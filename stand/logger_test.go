package stand

import (
	"context"
	"errors"
	"testing"

	"golang.org/x/exp/slog"
)

func TestInfo(t *testing.T) {
	logger := NewLogger(slog.LevelInfo)
	ctx := context.Background()
	ctx = context.WithValue(ctx, traceKey, "xxxxxx")
	ctx = context.WithValue(ctx, userKey, "wangbaoyi1")
	logger.InfoC(ctx, "info", "xxx")
	logger.ErrorC(ctx, errors.New("sql error "), "", "yyy")
}
