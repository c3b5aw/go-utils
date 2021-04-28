package log

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	} else {
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
