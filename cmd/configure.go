package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bcpitutor/tiki/appconfig"
	"github.com/spf13/cobra"
)

// type ServiceConfig struct {
// 	BaseUrl       string `yaml:"baseUrl"`
// 	ClientTimeout int    `yaml:"clientTimeOut"`
// }

// type BrowserConfig struct {
// 	PreferredBrowser string `yaml:"preferredBrowser"`
// }

// type TikiConfig struct {
// 	Browser BrowserConfig `yaml:"browser"`
// 	Service ServiceConfig `yaml:"service"`
// 	Debug   bool          `yaml:"debug"`
// 	AWSCli  string        `yaml:"awscli"`
// }

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Creates the configuration profile",
	Long:  ``,
	Run:   do_configure,
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

func do_configure(cmd *cobra.Command, args []string) {
	profile := rootCmd.Flag("profile").Value.String()
	fmt.Printf("Creating new config for profile: %s\n", profile)

	fmt.Printf("Preferred Browser: [Google Chrome] ")
	browserInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	browserVal := strings.Trim(browserInput, "\n")
	if browserVal == "" {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.browser", profile),
			"Google Chrome")
	} else {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.browser", profile),
			strings.Trim(browserInput, "\n"))
	}

	var urlVal string
	fmt.Print("Service baseURL: ")
	fmt.Scanf("%s", &urlVal)

	if urlVal == "" {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.baseurl", profile),
			"http://localhost:8080")
	} else {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.baseurl", profile),
			urlVal)
	}

	var debugVal string
	fmt.Print("Debug mode: [false]")
	fmt.Scanf("%s", &debugVal)

	if debugVal == "" {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.debug", profile),
			"false")
	} else {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.debug", profile),
			debugVal)
	}

	awscliPath, _ := exec.LookPath("aws")
	if awscliPath == "" {
		fmt.Println("It seems awscli is not installed on your system.")
		fmt.Println("If you plan to use aws cli features with tiki")
		fmt.Println("Please install it and re-try creating the configuration file")
		fmt.Println("Otherwise, just hit enter and continue")
	}

	var awsCliVal string
	fmt.Printf("AWS Cli Path: [%s]", awscliPath)
	fmt.Scanf("%s", &awsCliVal)

	if awsCliVal == "" {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.awscli", profile),
			awscliPath)
	} else {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.awscli", profile),
			awsCliVal)
	}

	// TODO: Ask if user wants to use kubectl subcommand
	var kubectlVal string
	fmt.Print("AWS Cli Path: [/usr/local/bin/kubectl]")
	fmt.Scanf("%s", &kubectlVal)

	if kubectlVal == "" {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.kubectl", profile),
			"/usr/local/bin/kubectl")
	} else {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.kubectl", profile),
			kubectlVal)
	}

	// TODO: Ask if user wants to use kubectl subcommand
	var defaultEKSCluster string
	fmt.Print("Default EKS Cluster to Use: []")
	fmt.Scanf("%s", &defaultEKSCluster)

	if defaultEKSCluster == "" {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.defaultEKSCluster", profile),
			"")
	} else {
		appconfig.AppConfig.ViperConf.Set(
			fmt.Sprintf("%s.defaultEKSCluster", profile),
			defaultEKSCluster)
	}

	err := appconfig.AppConfig.ViperConf.WriteConfig()
	if err != nil {
		fmt.Printf("Error writing config: %s\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Config profile %s created/updated\n", profile)
}
