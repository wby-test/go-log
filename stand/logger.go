package stand

import (
	"context"
	"io"

	"golang.org/x/exp/slog"
)

const (
	traceKey = "traceID"
	userKey  = "userID"
)

type Stand struct {
	logger *slog.Logger
}

func NewStandLogger(level slog.Level, w io.Writer) *Stand {
	// 替换掉 msg =
	opts := slog.HandlerOptions{
		AddSource: false,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				return slog.Attr{}
			}

			//if a.Key == slog.TimeKey {
			//	return slog.Attr{
			//		Key:   slog.TimeKey,
			//		Value: slog.StringValue(time.Now().Format("2006-01-02 15:04:05")),
			//	}
			//}

			return a
		},
	}

	textHandler := opts.NewTextHandler(w)
	logger := slog.New(textHandler)

	return &Stand{
		logger: logger,
	}
}

func (l Stand) Info(args ...any) {
	l.logger.Info("", args...)
}

func (l Stand) Warn(args ...any) {
	l.logger.Warn("", args...)
}

func (l Stand) Debug(args ...any) {
	l.logger.Debug("", args...)
}

func (l Stand) Error(err error, args ...any) {
	l.logger.Error("", err, args...)
}

func (l Stand) InfoC(ctx context.Context, args ...any) {
	ctxArgs := getContext(ctx)
	newArgs := setArgs(ctxArgs, args)
	l.Info(newArgs...)
}

func (l Stand) WarnC(ctx context.Context, args ...any) {
	ctxArgs := getContext(ctx)
	newArgs := setArgs(ctxArgs, args)
	l.Warn(newArgs...)
}

func (l Stand) DebugC(ctx context.Context, args ...any) {
	ctxArgs := getContext(ctx)
	newArgs := setArgs(ctxArgs, args)
	l.Debug(newArgs...)
}

func (l Stand) ErrorC(ctx context.Context, err error, args ...any) {
	ctxArgs := getContext(ctx)
	newArgs := setArgs(ctxArgs, args)
	l.Error(err, newArgs...)
}

func setArgs(ctxArgs, args []any) []any {
	newArgs := make([]any, len(args)+len(ctxArgs))
	copy(newArgs, ctxArgs)
	copy(newArgs[len(ctxArgs):], args)
	return newArgs
}

func getContext(ctx context.Context) (args []any) {
	reqNo := ctx.Value(traceKey)
	userID := ctx.Value(userKey)

	if reqNo == nil {
		reqNo = ""
	}

	if userID == nil {
		userID = ""
	}

	args = append(args, []any{"reqNo", reqNo, "userID", userID}...)
	return
}
