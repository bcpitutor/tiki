package service

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bcpitutor/tiki/backend"
	"github.com/bcpitutor/tiki/utils"
)

func RevokeSessionById(id string) map[string]any {
	postBody, err := utils.GetPostBody("revokeId", id)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodPost,
		"/session/revoke/id",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	return result
}

func RevokeSessionByEmail(email string) map[string]any {
	postBody, err := utils.GetPostBody("revokeEmail", email)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodPost,
		"/session/revoke/email",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	return result
}

func GetSessionList(reqSessionType string) map[string]any {
	parameters := url.Values{"sessionType": {reqSessionType}}

	URL, _ := url.Parse("/session/list")
	URL.RawQuery = parameters.Encode()
	url := URL.String()

	result, err := backend.ServiceCall(
		http.MethodGet,
		url,
		strings.NewReader(""),
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	return result
}
