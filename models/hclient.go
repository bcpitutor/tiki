package models

import (
	"net/http"
	"time"

	"github.com/bcpitutor/tiki/utils"
	"github.com/spf13/viper"
)

type HClient struct {
	c           *http.Client
	gToken      string
	rToken      string
	machineID   string
	machineInfo string
	userAgent   string
}

func NewHTTPClient(gToken string) *HClient {
	clientTimeOut := viper.GetInt("service.clientTimeout")
	timeOut := time.Second * time.Duration(clientTimeOut)
	var client = &http.Client{Timeout: timeOut}

	machineID, err := utils.GetMachineId()
	if err != nil {
		machineID = "PredefinedMachineId"
	}

	applicationID, err := utils.GetAppKey("tiki")
	if err != nil {
		applicationID = "PredefinedTikit"
	}

	userAgent := utils.GetUserAgent()

	return &HClient{
		c:           client,
		gToken:      gToken,
		machineID:   machineID,
		machineInfo: machineID + ":" + applicationID,
		userAgent:   userAgent,
	}
}

// Previous implemetation.
// It uses before migration of auth command to newAuth command.
func NewHTTPClientWithToken() *HClient {
	clientTimeOut := viper.GetInt("service.clientTimeout")
	timeOut := time.Second * time.Duration(clientTimeOut)
	var client = &http.Client{Timeout: timeOut}

	machineID, err := utils.GetMachineId()
	if err != nil {
		machineID = "PredefinedMachineId"
	}
	applicationID, err := utils.GetAppKey("tiki")
	if err != nil {
		applicationID = "PredefinedTikitool"
	}

	gToken, err := utils.GetCurrentToken()
	if err != nil {
		gToken = "token couldn't parsed"
	}

	rToken, err := utils.GetRefreshToken()
	if err != nil {
		rToken = "token couldn't parsed"
	}

	userAgent := utils.GetUserAgent()

	return &HClient{
		c:           client,
		gToken:      gToken,
		rToken:      rToken,
		machineID:   machineID,
		machineInfo: machineID + ":" + applicationID,
		userAgent:   userAgent,
	}

}

// Setters and Getters
func (hclient *HClient) GetToken() string {
	return hclient.gToken
}

func (hclient *HClient) SetToken(token string) {
	hclient.gToken = token
}

func (hclient *HClient) GetMachineID() string {
	return hclient.machineID
}

func (hclient *HClient) SetMachineID(machineID string) {
	hclient.machineID = machineID
}

func (hclient *HClient) GetMachineInfo() string {
	return hclient.machineInfo
}

func (hclient *HClient) SetMachineInfo(machineInfo string) {
	hclient.machineInfo = machineInfo
}

func (hclient *HClient) GetHttpClient() *http.Client {
	return hclient.c
}

func (hclient *HClient) GetUserAgent() string {
	return hclient.userAgent
}

func (hclient *HClient) SetUserAgent(userAgent string) {
	hclient.userAgent = userAgent
}
