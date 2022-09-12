package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	TokenFile = "tiki.token"
)

func EncryptToken(t string, k string) (string, error) {
	token := []byte(t)
	key := []byte(k)

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	result := gcm.Seal(nonce, nonce, token, nil)

	return string(result), nil
}

func DecryptToken(data []byte, key []byte) (string, error) {

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", err
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func DumpEncryptedToken(encryptedData []byte) error {
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
	fPath := filepath.Join(outdirpath, TokenFile)

	f, err := os.Create(fPath)
	if err != nil {
		return err
	}

	//fmt.Printf("Writing new token\n")

	if _, err := f.Write(encryptedData); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func ExportEncryptedToken() ([]byte, error) {
	homeFolder, err := HomeDirectory()
	if err != nil {
		return nil, err
	}

	outdirpath := fmt.Sprintf("%s/%s", homeFolder, ".tiki")
	fPath := filepath.Join(outdirpath, TokenFile)

	tokeEncyptedData, err := os.ReadFile(fPath)
	if err != nil {
		return nil, err
	}

	return tokeEncyptedData, nil
}
