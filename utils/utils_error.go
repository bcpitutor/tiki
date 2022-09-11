package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type errorJson struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ErrOutput(msg string, jsonFlag bool, exitApp bool) {
	if jsonFlag {
		ej := errorJson{
			Status:  "error",
			Message: msg,
		}
		b, _ := json.Marshal(ej)
		fmt.Printf("%s\n", string(b))
	} else {
		fmt.Printf("Error: %s\n", msg)
	}

	if exitApp {
		os.Exit(-1)
	}
}
