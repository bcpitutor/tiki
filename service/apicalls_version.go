package service

import (
	"net/http"
	"strings"

	"github.com/bcpitutor/tiki/backend"
)

func GetApiVersion() map[string]any {
	result, err := backend.ServiceCall(
		http.MethodGet,
		"/version",
		strings.NewReader(""),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

// func GetApiVersion() map[string]any {
// 	gToken, err := utils.GetCurrentToken()
// 	if err != nil {
// 		return map[string]any{"message": err, "status": "error"}
// 	}

// 	baseUrl := viper.GetString("service.baseUrl")
// 	versionUrl := baseUrl + "/version"

// 	hclient := models.NewHTTPClient(gToken)

// 	req, err := backend.NewTikiClientRequest(http.MethodGet, versionUrl, strings.NewReader(""))
// 	if err != nil {
// 		return map[string]any{"message": err, "status": "error"}
// 	}

// 	res, serviceErr := hclient.GetHttpClient().Do(req)
// 	if serviceErr != nil {
// 		return map[string]any{"message": serviceErr, "status": "error"}
// 	}

// 	defer res.Body.Close()

// 	var rbody map[string]interface{}

// 	if res.StatusCode >= 200 && res.StatusCode <= 299 {

// 		err := json.NewDecoder(res.Body).Decode(&rbody)
// 		if err != nil {
// 			if viper.GetBool("debug") {
// 				fmt.Printf("error while reading response body : %v", err)
// 			}
// 		}
// 		// TODO: send/return aws creds to the requester.
// 		return rbody
// 	} else {
// 		if viper.GetBool("debug") {
// 			fmt.Println("Service reponse failure : ", res.StatusCode)
// 		}
// 		return map[string]any{"message": err, "status": "error", "statusCode": res.StatusCode}
// 	}
// }
