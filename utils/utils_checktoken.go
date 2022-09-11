package utils

import (
	"fmt"
	"os"
)

func CheckRefreshToken(newToken any) {
	if newToken == nil {
		return
	}

	//fmt.Printf("Renewing token...\n")
	if newToken != "" {
		value := newToken.(string)
		encryptedData, err := EncryptToken(value, Enckey.GetEncKey())
		if err != nil {
			fmt.Printf("Encrypt operation failure : %v", err)
			os.Exit(1)
		}
		DumpEncryptedToken([]byte(encryptedData))
	}
}
