package log

import "io"

const (
	COLOR_RESET  string = "\033[0m"
	COLOR_GREEN  string = "\033[32m"
	COLOR_WHITE  string = "\033[37m"
	COLOR_RED    string = "\033[31m"
	COLOR_YELLOW string = "\033[33m"
	COLOR_PURPLE string = "\033[35m"
)

var (
	DEBUG bool = true
)

type Logger struct {
	Out io.Writer
}
