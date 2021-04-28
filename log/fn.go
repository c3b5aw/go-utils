package log

import (
	"fmt"
	"os"
)

func Printf(format string, v ...interface{}) {
	Debug(fmt.Sprintf(format, v))
}

func Leave() error {
	Info("Leaving worker safely...")
	os.Exit(0)
	return nil
}
