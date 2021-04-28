package log

import (
	"fmt"
	"os"
)

func Printf(format string, v ...interface{}) {
	Debug(fmt.Sprintf(format, v))
}

func Leave() {
	Info("Leaving worker safely...")
	os.Exit(0)
}
