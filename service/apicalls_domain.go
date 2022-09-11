package service

import (
	"net/http"
	"strings"

	"github.com/bcpitutor/tiki/backend"
	"github.com/bcpitutor/tiki/utils"
)

func DomainList() map[string]any {
	result, err := backend.ServiceCall(
		http.MethodGet,
		"/domain/list",
		strings.NewReader(""),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func DomainByName(domainPath string) map[string]any {
	postBody, err := utils.GetPostBody("domainPath", domainPath)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodPost,
		"/domain/get",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func CreateDomain(domainPath string, ownerGroup string) map[string]any {
	postBody, err := utils.GetPostBody("domainPath", domainPath, "ownerGroup", ownerGroup)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodPost,
		"/domain/create",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func DeleteDomain(domainPath string) map[string]any {
	postBody, err := utils.GetPostBody("domainPath", domainPath)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodDelete,
		"/domain/delete",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}
