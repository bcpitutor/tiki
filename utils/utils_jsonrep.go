package utils

import (
	"encoding/json"
	"fmt"
)

type TikitoolResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Count   string `json:"count"`
	Details string `json:"details"`
	Data    string `json:"data"`
}

func (response *TikitoolResponse) TikitoolResponseToJSON() string {
	resp, err := json.Marshal(&response)
	if err != nil {
		// TODO: becareful to loop and segmeatation fault.
		response.Details = fmt.Sprintf("%v", err)
	}
	return string(resp)
}

func TikiToolResponseAsIsToJSON(serviceResponse map[string]any) {
	j, err := json.MarshalIndent(serviceResponse, "", " ")
	if err != nil {
		fmt.Println("{\"error\":\"marshal error\"}")
	}
	fmt.Println(string(j))
}
