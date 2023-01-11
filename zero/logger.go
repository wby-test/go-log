package zero

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const _REQUEST_ID_KEY = "x-request-id"

func GetLogger(level zerolog.Level) zerolog.Logger {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    true,
	})
	log.Logger = log.Level(level)
	return log.Logger
}

func Debug(f string, args ...any) {
	log.Debug().Msgf(f, args...)
}

func Info(f string, args ...any) {
	log.Info().Msgf(f, args...)
}

func Warn(f string, args ...any) {
	log.Warn().Msgf(f, args...)
}

func Error(f string, args ...any) {
	log.Error().Msgf(f, args...)
}

func Fatal(f string, args ...any) {
	log.Fatal().Msgf(f, args...)
}

func Panic(f string, args ...any) {
	log.Panic().Msgf(f, args...)
}

func DebugC(c context.Context, f string, args ...any) {
	writer(c, zerolog.DebugLevel, f, args...)
}

func InfoC(c context.Context, f string, args ...any) {
	writer(c, zerolog.InfoLevel, f, args...)
}

func WarnC(c context.Context, f string, args ...any) {
	writer(c, zerolog.WarnLevel, f, args...)
}

func ErrorC(c context.Context, f string, args ...any) {
	writer(c, zerolog.ErrorLevel, f, args...)
}

func FatalC(c context.Context, f string, args ...any) {
	writer(c, zerolog.FatalLevel, f, args...)
}

func PaincC(c context.Context, f string, args ...any) {
	writer(c, zerolog.PanicLevel, f, args...)
}

func ErrorT(c context.Context, level zerolog.Level, f string, args ...any) {
	log.WithLevel(level).Str(_REQUEST_ID_KEY, "xxxx").Msgf(f, args...)
}

func writer(c context.Context, level zerolog.Level, f string, args ...any) {
	reqId, ok := "test", true
	if !ok {
		log.Warn().Msg("log writer with trace err")
	}

	log.WithLevel(level).Str(_REQUEST_ID_KEY, reqId).Msgf(f, args...)
}
