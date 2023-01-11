package logger

import "context"

type Logger interface {
	Info(args ...any)
	Warn(args ...any)
	Debug(args ...any)
	Error(err error, args ...any)
	InfoC(ctx context.Context, args ...any)
	WarnC(ctx context.Context, args ...any)
	DebugC(ctx context.Context, args ...any)
	ErrorC(ctx context.Context, err error, args ...any)
}
