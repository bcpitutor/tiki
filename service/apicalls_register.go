package service

// func CreateRegistrant(registrantInfo map[string]interface{}) map[string]interface{} {

// 	var result map[string]interface{}
// 	hclient := models.NewHTTPClientWithToken()
// 	baseUrl := viper.GetString("service.baseUrl")
// 	registerCreateUrl := baseUrl + "/register/create"

// 	postBody, err := json.Marshal(registrantInfo)
// 	if err != nil {
// 		errMessage := map[string]interface{}{
// 			"message": err,
// 			"status":  "error",
// 		}
// 		return errMessage
// 	}
// 	// req, err := http.NewRequest(http.MethodPost, registerCreateUrl, strings.NewReader(string(postBody)))
// 	req, err := backend.NewTikiClientRequest(http.MethodPost, registerCreateUrl, strings.NewReader(string(postBody)))
// 	if err != nil {
// 		errMessage := map[string]interface{}{
// 			"message": err,
// 			"status":  "error",
// 		}
// 		// utils.TikiToolResponseAsIsToJSON(errMessage)
// 		return errMessage
// 	}

// 	// baerer := "Baerer " + hclient.GetToken()
// 	// req.Header.Set("Authorization", baerer)

// 	res, serviceErr := hclient.GetHttpClient().Do(req)
// 	if serviceErr != nil {
// 		errMessage := map[string]interface{}{
// 			"details": serviceErr,
// 			"message": "service response error",
// 			"status":  "error",
// 		}

// 		return errMessage
// 	}

// 	defer res.Body.Close()

// 	err = json.NewDecoder(res.Body).Decode(&result)
// 	if err != nil {
// 		errMessage := map[string]interface{}{
// 			"message": err,
// 			"status":  "error",
// 		}
// 		return errMessage
// 	}

// 	return result
// }
