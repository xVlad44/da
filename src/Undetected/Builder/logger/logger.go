package logger

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	info  = color.New(color.FgGreen).PrintlnFunc()
	err   = color.New(color.FgRed).PrintlnFunc()
	fatal = color.New(color.BgRed).PrintlnFunc()
)

func Info(contents ...interface{}) {
	content := formatMessage(contents...)
	info("~Pixon > ", content)
}

func Error(contents ...interface{}) {
	content := formatMessage(contents...)
	err("~Pixon > ", content)
}

func Fatal(contents ...interface{}) {
	content := formatMessage(contents...)
	fatal("~Pixon > ", content)
}

func formatMessage(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = strings.TrimRight(msg, " \n\r")
	return msg
}
