package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger(level zerolog.Level) ZeroLogger {
	logger := log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    true,
	})
	logger = logger.Level(level)
	return ZeroLogger{logger: logger}
}

func (l ZeroLogger) Debug(args ...any) {
	l.logger.Debug().Msgf(getConnect(args))
}

func (l ZeroLogger) Info(args ...any) {
	l.logger.Info().Msgf(getConnect(args))
}

func (l ZeroLogger) Warn(args ...any) {
	l.logger.Warn().Msgf(getConnect(args))
}

func (l ZeroLogger) Error(args ...any) {
	l.logger.Error().Msgf(getConnect(args))
}

func (l ZeroLogger) Fatal(args ...any) {
	l.logger.Fatal().Msgf(getConnect(args))
}

func (l ZeroLogger) Panic(args ...any) {
	l.logger.Panic().Msgf(getConnect(args))
}

func (l ZeroLogger) DebugC(c context.Context, args ...any) {
	l.writer(c, zerolog.DebugLevel, args...)
}

func (l ZeroLogger) InfoC(c context.Context, args ...any) {
	l.writer(c, zerolog.InfoLevel, args...)
}

func (l ZeroLogger) WarnC(c context.Context, args ...any) {
	l.writer(c, zerolog.WarnLevel, args...)
}

func (l ZeroLogger) ErrorC(c context.Context, args ...any) {
	l.writer(c, zerolog.ErrorLevel, args...)
}

func (l ZeroLogger) FatalC(c context.Context, args ...any) {
	l.writer(c, zerolog.FatalLevel, args...)
}

func (l ZeroLogger) PanicC(c context.Context, args ...any) {
	l.writer(c, zerolog.PanicLevel, args...)
}

func (l ZeroLogger) ErrorT(c context.Context, level zerolog.Level, args ...any) {
	l.logger.WithLevel(level).Str(traceKey, "xxxx").Msgf(getConnect(args))
}

func (l ZeroLogger) writer(c context.Context, level zerolog.Level, args ...any) {
	reqId, ok := "test", true
	if !ok {
		l.logger.Warn().Msg("log writer with trace err")
	}

	l.logger.WithLevel(level).Str(traceKey, reqId).Msgf(getConnect(args))
}

func getConnect(args ...any) (f string, newArgs []any) {
	for k, v := range args {
		if k == 0 {
			f = v.(string)
			continue
		}
		newArgs = append(newArgs, v)
	}

	return
}
