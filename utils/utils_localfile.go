package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func DumpEncryptedDataToFile(encryptedData []byte, outputFile string) error {
	homeFolder, err := HomeDirectory()
	if err != nil {
		return err
	}
	outdirpath := fmt.Sprintf("%s/%s", homeFolder, ".tiki")

	fMode := fs.FileMode(uint32(0700))
	err = os.MkdirAll(outdirpath, fs.FileMode(fMode))
	if err != nil {
		return err
	}
	fPath := filepath.Join(outdirpath, outputFile)

	f, err := os.Create(fPath)
	if err != nil {
		return err
	}

	if _, err := f.Write(encryptedData); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
