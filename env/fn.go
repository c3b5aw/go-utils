package env

import (
	"os"
	"strings"
)

func Bool(k string, def bool) bool {
	if strings.ToLower(os.Getenv(k)) == "true" {
		return true
	}
	return def
}

func String(k, def string) string {
	if os.Getenv(k) != "" {
		return os.Getenv(k)
	}
	return def
}
