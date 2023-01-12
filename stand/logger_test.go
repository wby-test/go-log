package stand

import (
	"bytes"
	"context"
	"errors"
	"os"
	"regexp"
	"strings"
	"testing"

	"golang.org/x/exp/slog"
)

const timeRE = `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}(Z|[+-]\d{2}:\d{2})`

// 表格驱动测试
func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStandLogger(slog.LevelDebug, &buf)
	check := func(want string) {
		t.Helper()
		if want != "" {
			want = "time=" + timeRE + " " + want
		}
		checkLogOutput(t, buf.String(), want)
		buf.Reset()
	}
	var tests = []struct {
		name  string
		input []any
		want  string
	}{
		{"info", []any{"a", 1, "b", 2}, "level=INFO a=1 b=2"},
		{"debug", []any{"a", 1, "b", 2}, "level=DEBUG a=1 b=2"},
		{"warn", []any{"a", 1, "b", 2}, "level=WARN a=1 b=2"},
		{"error", []any{"a", 1, "b", 2}, "level=ERROR a=1 b=2 err=test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "error" {
				logger.Error(errors.New("test"), tt.input...)
				check(tt.want)
			} else {
				logger.Info(tt.input...)
				check(tt.want)
			}
		})
	}
}

func TestLogOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStandLogger(slog.LevelDebug, &buf)
	check := func(want string) {
		t.Helper()
		if want != "" {
			want = "time=" + timeRE + " " + want
		}
		checkLogOutput(t, buf.String(), want)
		buf.Reset()
	}

	logger.Info("a", 1, "b", 2)
	check(`level=INFO a=1 b=2`)
	logger.Error(errors.New("test"), "a", 1, "b", 2)
	check(`level=ERROR a=1 b=2 err=test`)
}

func getTestContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, traceKey, "xxxxxx")
	ctx = context.WithValue(ctx, userKey, "wangbaoyi1")
	return ctx
}

func checkLogOutput(t *testing.T, got, wantRegexp string) {
	t.Helper()
	got = clean(got)
	wantRegexp = "^" + wantRegexp + "$"
	matched, err := regexp.MatchString(wantRegexp, got)
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("\ngot  %s\nwant %s", got, wantRegexp)
	}
}

// clean prepares log output for comparison.
func clean(s string) string {
	if len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	return strings.ReplaceAll(s, "\n", "~")
}

func TestStand_Info(t *testing.T) {
	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Stand{
				logger: tt.fields.logger,
			}
			l.Info(tt.args.args...)
		})
	}
}

func TestStand_Warn(t *testing.T) {
	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		args []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"test1", fields{logger: NewStandLogger(slog.LevelDebug, os.Stdout).logger}, args{
			[]any{"a", 1, "b", 2},
		}},
		{"test2", fields{logger: NewStandLogger(slog.LevelDebug, os.Stdout).logger}, args{
			[]any{"a", 1, "b", 2},
		}},
		{"test3", fields{logger: NewStandLogger(slog.LevelDebug, os.Stdout).logger}, args{
			[]any{"a", 1, "b", 2},
		}},
		{"test4", fields{logger: NewStandLogger(slog.LevelDebug, os.Stdout).logger}, args{
			[]any{"a", 1, "b", 2},
		}},
		{"test5", fields{logger: NewStandLogger(slog.LevelDebug, os.Stdout).logger}, args{
			[]any{"a", 1, "b", 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Stand{
				logger: tt.fields.logger,
			}
			l.Warn(tt.args.args...)
		})
	}
}
