package log

import "io"

const (
	COLOR_RESET  = "\033[0m"
	COLOR_GREEN  = "\033[32m"
	COLOR_WHITE  = "\033[37m"
	COLOR_RED    = "\033[31m"
	COLOR_YELLOW = "\033[33m"
	COLOR_PURPLE = "\033[35m"
)

type Logger struct {
	Out io.Writer
}
