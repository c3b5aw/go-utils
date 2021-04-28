package log

import (
	"fmt"
	"os"
	"time"
)

var (
	Std = New(os.Stderr)
)

func print(header, msg, color string, params []string) (int, error) {
	var buf []byte

	timeHeader := time.Now().Format("15:04:05.000000")

	formattedMsg := fmt.Sprintf("\t[%s%s%s] ", color, header, COLOR_RESET)
	// Set sender if any
	if len(params) >= 1 {
		formattedMsg = formattedMsg + fmt.Sprintf("[%s] ", params[0])
	}
	formattedMsg = formattedMsg + fmt.Sprintf("%s\n", msg)

	buf = append(buf, timeHeader+formattedMsg...)
	return Std.Out.Write(buf)
}
