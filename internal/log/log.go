package log

import (
	hclog "github.com/hashicorp/go-hclog"
)

func Info(msg string, args ...interface{}) {
	hclog.Default().Info(msg, args...)
}

func Error(msg string, args ...interface{}) {
	hclog.Default().Error(msg, args...)
}

func Fmt(str string, args ...interface{}) hclog.Format {
	return hclog.Fmt(str, args...)
}
