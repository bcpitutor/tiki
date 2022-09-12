package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/denisbrodbeck/machineid"
)

func GetPostBody(postArgs ...string) ([]byte, error) {
	// Create postBodyData out of the postArgs as a map[string]string
	postBodyData := map[string]string{}
	for i := 0; i < len(postArgs); i += 2 {
		postBodyData[postArgs[i]] = postArgs[i+1]
	}

	postBody, err := json.Marshal(postBodyData)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshalling postBodyData: %v", err)
	}

	return postBody, nil
}

func GetRefreshToken() (string, error) {
	homeFolder, err := HomeDirectory()
	if err != nil {
		return "", nil
	}

	encFile := "tiki.rt.data"
	fPath := filepath.Join(homeFolder, "/.tiki", encFile)

	localTokenStore, err := os.ReadFile(fPath)
	if err != nil {
		return "", err
	}

	eKey := Enckey.GetEncKey()
	rToken, err := DecryptToken(localTokenStore, []byte(eKey))
	if err != nil {
		return "", err
	}

	return rToken, nil
}

func GetCurrentToken() (string, error) {
	homeFolder, err := HomeDirectory()
	if err != nil {
		return "", nil
	}

	encFile := TokenFile
	fPath := filepath.Join(homeFolder, "/.tiki", encFile) // TODO: update directory

	localTokenStore, err := os.ReadFile(fPath)
	if err != nil {
		return "", err
	}

	eKey := Enckey.GetEncKey()
	gToken, err := DecryptToken(localTokenStore, []byte(eKey))
	if err != nil {
		return "", err
	}

	return gToken, nil
}

func GetUserAgent() string {
	currentVersion := ""
	return "Tiki/" + currentVersion
}

func GetMachineId() (string, error) {
	id, err := machineid.ID()
	if err != nil {
		return "unavailable", err
	}
	return id, nil
}
