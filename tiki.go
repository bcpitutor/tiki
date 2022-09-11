package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcpitutor/tiki/cmd"
	"github.com/bcpitutor/tiki/utils"
)

func init() {
	mId, err := utils.GetMachineId()
	if err != nil {
		fmt.Printf("Cannot detect Machine ID, quitting. Error: %v\n", err)
		os.Exit(-1)
	}

	mIdTrimmed := strings.Replace(mId, "-", "", 4)

	if mIdTrimmed != "" {
		utils.Enckey.SetEncKey(mIdTrimmed)
	} else {
		// TODO: Check if this is a valid operation?
		//utils.Enckey.SetEncKey("01234567012345670123456776543210")
		fmt.Printf("Cannot generate Machine ID, quitting.\n")
		os.Exit(1)
	}
}

func main() {
	cmd.Execute()
}
