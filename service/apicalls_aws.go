package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bcpitutor/tiki/backend"
)

func ListAWSRoles() map[string]any {
	result, err := backend.ServiceCall(
		http.MethodGet,
		"/aws/roles/list",
		strings.NewReader(""),
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	return result
}
