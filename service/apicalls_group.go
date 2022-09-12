package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bcpitutor/tiki/backend"
	"github.com/bcpitutor/tiki/utils"
)

// GET http(s)://baseURL/group

func GetGroupList() map[string]any {
	result, err := backend.ServiceCall(
		http.MethodGet,
		"/group",
		strings.NewReader(""),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

// GET http(s)://baseURL/group/groupName
func FindGroupByName(groupName string) map[string]any {
	if groupName == "" {
		return map[string]any{"message": "groupName is missing", "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodGet,
		fmt.Sprintf("/group/%s", groupName),
		strings.NewReader(""),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func CreateGroup(dataFilePath string) map[string]any {
	if dataFilePath == "" {
		return map[string]any{"message": "dataFilePath is missing", "status": "error"}
	}

	postBody, _ := os.ReadFile(dataFilePath)
	result, err := backend.ServiceCall(
		http.MethodPost,
		"/group/create",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func DeleteGroup(groupName string) map[string]any {
	if groupName == "" {
		return map[string]any{"message": "groupName is missing", "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodDelete,
		fmt.Sprintf("/group/delete/%s", groupName),
		strings.NewReader(""),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func AddMemberToGroup(groupName string, memberEmail string) map[string]any {
	postBody, _ := utils.GetPostBody(
		"newMemberEmail", memberEmail,
	)

	result, err := backend.ServiceCall(
		http.MethodPost,
		fmt.Sprintf("/group/addmember/%s", groupName),
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func DeleteMemberFromGroup(groupName string, memberEmail string) map[string]any {
	postBody, _ := utils.GetPostBody(
		"deleteMemberEmail", memberEmail,
	)
	result, err := backend.ServiceCall(
		http.MethodPost,
		fmt.Sprintf("/group/delmember/%s", groupName),
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}
