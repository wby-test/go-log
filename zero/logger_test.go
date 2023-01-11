// Package logger provides ...
package zero

import (
	"context"
	"testing"

	//"ai.jdt.com/ai-platform/operationcenter/internal/pkg/logger"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func TestTraceLog(*testing.T) {
	ctx := context.Background()
	id := uuid.New().String()
	ctx = context.WithValue(ctx, "x-request-id", id)

	InfoC(ctx, "this is test")
}

func TestWarn(t *testing.T) {
	GetLogger(zerolog.ErrorLevel)
	Warn("test warn %s", "tttt")
	Info("test info %s", "xxxxx")
	Error("test err %s", "ppppp")
	ErrorT(nil, zerolog.InfoLevel, "writ %s", "sss")
}
