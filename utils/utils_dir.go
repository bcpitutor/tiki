package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"runtime"

	"github.com/denisbrodbeck/machineid"
)

func HomeDirectory() (string, error) {
	homeEnvValue := ""

	switch runtime.GOOS {
	case "darwin", "linux":
		homeEnvValue = "HOME"
	// TODO: Add Windows support maybe?
	default:
		return "", errors.New("unsupported platform")
	}

	home := os.Getenv(homeEnvValue)
	return home, nil
}

func InitalizeTikiDirectory() string {
	home, err := HomeDirectory()
	if err != nil {
		fmt.Printf("Cannot detect home folder: %s\n", err)
		os.Exit(-1)
	}

	outdirpath := fmt.Sprintf("%s%s", home, "/.tiki")
	err = os.MkdirAll(
		outdirpath,
		fs.FileMode(fs.FileMode(uint32(0700))),
	)
	if err != nil {
		fmt.Printf("Tiki config directory does not exist and failed to create one: %s\n", err)
		os.Exit(-1)
	}

	return outdirpath
}

func GetAppKey(appName string) (string, error) {
	id, err := machineid.ProtectedID(appName)
	if err != nil {
		return "unavailable", err
	}
	return id, nil
}
