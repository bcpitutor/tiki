package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func OpenBrowser(preferredBrowser string, url string) bool {
	if preferredBrowser == "" {
		preferredBrowser = "Google Chrome"
	}

	var args []string
	aws_env := os.Getenv("AWS_ENV")
	switch runtime.GOOS {
	case "darwin":
		if aws_env != "" {
			args = []string{
				"open",
				"-na",
				preferredBrowser,
				"--args",
				"--profile-directory=" + aws_env,
				"--new-window"}
		} else {
			args = []string{
				"open",
				"-na",
				preferredBrowser,
				"--args"}
		}
	case "windows":
		// TODO: This would require some more work to support multiple browsers
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}

	cmd := exec.Command(args[0], append(args[1:], url)...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}
