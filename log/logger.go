package log

import (
	"io"
)

func New(out io.Writer) *Logger {
	return &Logger{Out: out}
}

// - Methods

func Debug(msg string, params ...string) {
	print("DEBUG", msg, COLOR_GREEN, params)
}

func Info(msg string, params ...string) {
	print("INFOS", msg, COLOR_WHITE, params)
}

func Success(msg string, params ...string) {
	print("SUCCS", msg, COLOR_YELLOW, params)
}

func Warning(msg string, params ...string) {
	print("WARN!", msg, COLOR_PURPLE, params)
}

func Error(msg string, params ...string) {
	print("ERROR", msg, COLOR_RED, params)
}
