package service

import (
	"net/http"
	"strings"

	"github.com/bcpitutor/tiki/backend"
	"github.com/bcpitutor/tiki/utils"
)

func GetBannedUserList() map[string]any {
	result, err := backend.ServiceCall(
		http.MethodGet,
		"/user/banlist",
		strings.NewReader(""),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func BanUser(email string) map[string]any {
	postBody, err := utils.GetPostBody("userEmail", email)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodPost,
		"/user/ban",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	return result
}

func UnBanUser(email string) map[string]any {
	postBody, err := utils.GetPostBody("userEmail", email)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	result, err := backend.ServiceCall(
		http.MethodDelete,
		"/user/unban",
		strings.NewReader(string(postBody)),
	)
	if err != nil {
		return map[string]any{"message": err, "status": "error"}
	}

	// var result map[string]any

	// hclient := models.NewHTTPClientWithToken()
	// baseURL := viper.GetString("service.baseUrl")
	// sessListUrl := baseURL + "/user/unban"

	// parameters := url.Values{}
	// parameters.Add("userEmail", email)

	// URL, err := url.Parse(sessListUrl)
	// URL.RawQuery = parameters.Encode()
	// url := URL.String()

	// req, err := backend.NewTikiClientRequest(
	// 	http.MethodDelete,
	// 	url,
	// 	strings.NewReader(""),
	// )

	// if err != nil {
	// 	errMessage := map[string]any{
	// 		"message": err,
	// 		"status":  "error",
	// 	}
	// 	return errMessage
	// }

	// res, serviceErr := hclient.GetHttpClient().Do(req)
	// if serviceErr != nil {
	// 	errMessage := map[string]any{
	// 		"details": serviceErr,
	// 		"message": "service response error",
	// 		"status":  "error",
	// 	}
	// 	return errMessage
	// }

	// defer res.Body.Close()

	// err = json.NewDecoder(res.Body).Decode(&result)
	// if err != nil {
	// 	errMessage := map[string]any{
	// 		"message": err,
	// 		"status":  "error",
	// 	}
	// 	return errMessage
	// }

	return result
}
