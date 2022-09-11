package utils

// TODO:  cleanup

// func getConfigFile(tikidir string) string {
// 	configFiles := []string{}

// 	_, err := os.Stat(tikidir + "/tiki.ini")
// 	if err != nil {
// 		files, err := ioutil.ReadDir(tikidir)
// 		if err != nil {
// 			fmt.Printf("Cannot read directory: %s\n", err)
// 			os.Exit(-1)
// 		}

// 		for _, file := range files {
// 			if file.IsDir() {
// 				continue
// 			}
// 			configFiles = append(configFiles, file.Name())
// 		}

// 		if len(configFiles) == 0 {
// 			return ""
// 		}

// 		fmt.Printf("Found multiple config files: %s\n", configFiles)
// 		os.Exit(0)
// 	}

// 	return tikidir + "/tiki.ini"

// 	// if _, err := os.Stat(tikidir + "/tiki.ini"); err == nil {
// 	// 	return tikidir + "/tiki.ini"
// 	// }

// 	// files, err := ioutil.ReadDir(tikidir)
// 	// if err != nil {
// 	// 	fmt.Printf("Cannot read directory: %s\n", err)
// 	// 	os.Exit(-1)
// 	// }

// 	// for _, file := range files {
// 	// 	if file.IsDir() {
// 	// 		continue
// 	// 	}
// 	// 	if filepath.Ext(file.Name()) == ".ini" {
// 	// 		configFiles = append(configFiles, file.Name())
// 	// 	}
// 	// }

// 	// fmt.Printf("%+v\n", configFiles)

// 	// os.Exit(-1)
// 	// return ""
// }

// func ReadConfigFile(tikidir string) {
// 	fileloc := fmt.Sprintf("%s%s", tikidir, "/tiki.ini")
// 	_, err := os.Stat(fileloc)
// 	if err != nil {
// 		fmt.Printf("Cannot read config file: %s\n", err)
// 		os.Exit(-1)
// 	}

// 	appconfig.InitConfig(fileloc)

// }
