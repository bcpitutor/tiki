package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bcpitutor/tiki/appconfig"
	"github.com/bcpitutor/tiki/models"
	"github.com/bcpitutor/tiki/utils"
	"github.com/google/uuid"
)

func ServiceCall(method string, url string, body io.Reader) (map[string]any, error) {
	hclient := models.NewHTTPClientWithToken()

	profile := appconfig.AppConfig.SelectedProfile

	urlKey := fmt.Sprintf("%s.baseurl", profile)
	baseURL := appconfig.AppConfig.ViperConf.GetString(urlKey)

	serviceURL := baseURL + url
	var result map[string]any

	//fmt.Printf("ServiceCall: %s %s\n", method, serviceURL)

	req, err := NewTikiClientRequest(method, serviceURL, body)
	if err != nil {
		return nil, fmt.Errorf("unexpected error while creating request: %+v", err)
	}

	// TODO: REMOVE
	if url == "/renewToken" {
		req.Header.Set("Hede", "hodo")
	}

	res, serviceErr := hclient.GetHttpClient().Do(req)
	if serviceErr != nil {
		return nil, fmt.Errorf("unexpected error using HTTP Client: %+v", serviceErr)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		if res.StatusCode == 401 {
			return nil, fmt.Errorf("attempt to unauthorized access. please login again, your session could be expired")
		}
		return nil, fmt.Errorf("unexpected service response error: %s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("response decoding error while communicating with service: %+v", err)
	}

	return result, nil
}

func NewTikiClientRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	reqId, _ := uuid.NewRandom()
	req.Header.Set("x-tsreq-id", reqId.String())

	gToken, err := utils.GetCurrentToken()
	if err != nil {
		gToken = "token couldn't parsed"
	}

	rToken, err := utils.GetRefreshToken()
	if err != nil {
		rToken = "token couldn't parsed"
	}

	machineID, err := utils.GetMachineId()
	if err != nil {
		machineID = "PredefinedMachineId"
	}
	applicationID, err := utils.GetAppKey("tiki")
	if err != nil {
		applicationID = "PredefinedTiki"
	}

	userAgent := utils.GetUserAgent()

	baerer := "Baerer " + gToken
	req.Header.Set("Authorization", baerer)
	req.Header.Set("rToken", rToken)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("x-tiki-id", machineID+":"+applicationID)

	return req, nil
}
