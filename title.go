package title

import (
	"os"
	"os/exec"
	"runtime"
)

// Set a title on the terminal for macos and windows
func Set(title string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "title", title)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return
		}
	} else {
		cmd := exec.Command("echo", "-n", "\033]0;"+title+"\007")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			return
		}
		err = cmd.Wait()
		if err != nil {
			return
		}
	}
}
