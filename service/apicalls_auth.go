package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bcpitutor/tiki/backend"
)

func AuthRenew() map[string]any {
	result, err := backend.ServiceCall(
		http.MethodPost,
		"/renewToken",
		strings.NewReader(""),
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	return result
}
