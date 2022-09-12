package service

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bcpitutor/tiki/backend"
	"github.com/bcpitutor/tiki/utils"
)

// status: fixed
func GetTicketList() (map[string]any, error) {
	result, err := backend.ServiceCall(
		http.MethodGet,
		"/ticket/list",
		strings.NewReader(""),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ObtainTicket(ticketPath string) (map[string]any, error) {
	postBody, _ := utils.GetPostBody("ticketPath", ticketPath)
	result, err := backend.ServiceCall(
		http.MethodPost,
		"/ticket/obtain",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func SetTicketSecret(ticketPath string, secretData string) (map[string]any, error) {
	postBody, _ := utils.GetPostBody("ticketPath", ticketPath, "secretData", secretData)
	result, err := backend.ServiceCall(
		http.MethodPost,
		"/ticket/secret/set",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTicketByPath(path string) (map[string]any, error) {
	postBody, _ := utils.GetPostBody("ticketPath", path)
	result, err := backend.ServiceCall(
		http.MethodPost,
		"/ticket/get",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteTicket(ticketPath string) map[string]any {
	postBody, _ := utils.GetPostBody("ticketPath", ticketPath)
	result, err := backend.ServiceCall(
		http.MethodDelete,
		"/ticket/delete",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "error"}
	}

	return result
}

func GetTicketSecret(ticketPath string) (map[string]any, error) {
	postBody, _ := utils.GetPostBody("ticketPath", ticketPath)
	result, err := backend.ServiceCall(
		http.MethodPost,
		"/ticket/secret/get",
		bytes.NewBuffer(postBody),
	)
	if err != nil {
		return map[string]any{"message": err.Error(), "status": "error"}, fmt.Errorf("%+v", err)
	}

	return result, nil
}

func CreateTicket(ticketDataFilePath string) (map[string]any, error) {
	fileData, err := os.ReadFile(ticketDataFilePath)
	if err != nil {
		fmt.Printf("Data file cannot be found: %v", ticketDataFilePath)
		os.Exit(0)
	}

	result, err := backend.ServiceCall(
		http.MethodPost,
		"/ticket/create",
		bytes.NewBuffer(fileData),
	)
	if err != nil {
		return map[string]any{
			"message": err.Error(),
			"status":  "error",
		}, fmt.Errorf("%+v", err)
	}

	return result, nil
}
