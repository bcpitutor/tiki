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
		utils.ErrOutput(
			fmt.Sprintf("Cannot detect Machine ID, quitting. Error: %v\n", err.Error()),
			false,
			true)
		os.Exit(-1)
	}

	mIdTrimmed := strings.Replace(mId, "-", "", 4)

	if mIdTrimmed != "" {
		utils.Enckey.SetEncKey(mIdTrimmed)
	} else {
		// TODO: Check if this is a valid operation?
		//utils.Enckey.SetEncKey("01234567012345670123456776543210")
		utils.ErrOutput(
			"Cannot generate Machine ID, quitting",
			false,
			true)
		os.Exit(-1)
	}
}

func main() {
	cmd.Execute()
}
