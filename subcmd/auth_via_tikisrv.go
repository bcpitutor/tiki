package subcmd

// func GoogleAuthenticationViaTikiServer(profile string, cmd *cobra.Command, args []string) {
// 	port, _ := cmd.Flags().GetString("listenerPort")
// 	host := "localhost"

// 	var idToken string
// 	var respData map[string]any

// 	wg := sync.WaitGroup{}
// 	wg.Add(1)

// 	r := gin.Default()
// 	r.Use(
// 		cors.New(cors.Config{
// 			AllowAllOrigins:  true,
// 			AllowMethods:     []string{"GET", "POST"},
// 			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
// 			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
// 			AllowCredentials: true,
// 			MaxAge:           12 * time.Hour,
// 		}),
// 	)

// 	r.POST("/receiveToken", func(c *gin.Context) {
// 		if err := c.ShouldBindJSON(&respData); err != nil {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"message": "marshal error",
// 				"details": err,
// 			})
// 			return
// 		}

// 		idToken = respData["token"].(string)

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "ok",
// 		})

// 		defer wg.Done()
// 	})

// 	go r.Run(host + ":" + port)

// 	machineId, err := utils.GetMachineId()
// 	if err != nil {
// 		errMessage := map[string]any{
// 			"status":  "error",
// 			"message": "MachineID error",
// 			"details": fmt.Sprintf("%v", err),
// 		}

// 		utils.TikiToolResponseAsIsToJSON(errMessage)
// 		return
// 	}

// 	urlKey := fmt.Sprintf("%s.baseurl", profile)
// 	baseURL := appconfig.AppConfig.ViperConf.GetString(urlKey)

// 	url := baseURL + "/auth/" + machineId + "/" + port

// 	browserKey := fmt.Sprintf("%s.browser", profile)
// 	browser := appconfig.AppConfig.ViperConf.GetString(browserKey)
// 	utils.OpenBrowser(viper.GetString(browser), url)
// 	wg.Wait()

// 	if idToken == "" {
// 		fmt.Printf("Something wrong with your credentials\n")
// 	}

// 	// if viper.GetBool("debug") {
// 	// 	zlogger.Log.Sugar().Infof("idToken [%v]", idToken)
// 	// }

// 	encryptedData, err := utils.EncryptToken(idToken, utils.Enckey.GetEncKey())
// 	if err != nil {
// 		fmt.Printf("Encrypt operation failure : %v", err)
// 		fmt.Printf("Encrypted binary : %v", encryptedData)
// 		return
// 	}

// 	if err := utils.DumpEncryptedToken([]byte(encryptedData)); err != nil {
// 		fmt.Printf("Dump encrypt token opertaion failure : %v", err)
// 	}

// 	if respData["rToken"] != nil {
// 		rToken, err := utils.EncryptToken(respData["rToken"].(string), utils.Enckey.GetEncKey())
// 		if err != nil {
// 			fmt.Printf("Encrypt operation failure. [%v]", err)
// 			return
// 		}

// 		if err := utils.DumpEncryptedDataToFile([]byte(rToken), "tiki.rt.data"); err != nil {
// 			fmt.Printf("Encrypt operation during data dump. Try again. [%v]", err)
// 			return
// 		}
// 	}

// 	if respData["clientID"] != nil {
// 		clientID, err := utils.EncryptToken(respData["clientID"].(string), utils.Enckey.GetEncKey())
// 		if err != nil {
// 			fmt.Printf("Encrypt operation failure. Try again. [%v]", err)
// 			return
// 		}

// 		if err := utils.DumpEncryptedDataToFile([]byte(clientID), "tiki.c1.data"); err != nil {
// 			fmt.Printf("Encrypt operation during data dump. Try again. [%v]", err)
// 			return
// 		}
// 	}

// 	if respData["clientSecret"] != nil {
// 		clientSecret, err := utils.EncryptToken(respData["clientSecret"].(string), utils.Enckey.GetEncKey())
// 		if err != nil {
// 			fmt.Printf("Encrypt operation failure.. Try again. [%v]", err)
// 			return
// 		}
// 		if err := utils.DumpEncryptedDataToFile([]byte(clientSecret), "tiki.c2.data"); err != nil {
// 			fmt.Printf("Encrypt operation during data dump. Try again. [%v]", err)
// 			return
// 		}
// 	}

// 	if respData["clientSecret"] == nil || respData["clientID"] == nil || respData["rToken"] == nil {
// 		fmt.Println("Info: some of responses are empty. Your session can not be refrehed.")
// 	}

// 	fmt.Printf("\n\n")

// 	fmt.Println("---------------------------iTutor-------------------------------------")
// 	fmt.Println("At this moment, you should see a browser window opened. Please login to your ")
// 	fmt.Println("Google account. If not, that would mean that you have either not installed")
// 	fmt.Println("default browser Google Chrome or you specified a wrong browser in the config")
// 	fmt.Println("If it still won't open, please copy the url below and paste into a browser")
// 	fmt.Println("manually")
// 	fmt.Println(url)
// 	fmt.Println("#---------------------------------------------------------------------#")

// }
