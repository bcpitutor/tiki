package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/bcpitutor/tiki/appconfig"
	"github.com/bcpitutor/tiki/models"
	"github.com/bcpitutor/tiki/service"
	"github.com/bcpitutor/tiki/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	consoleSessDuration = time.Duration(3) * time.Hour // 3 Hours
	awsFederationURL    = "https://signin.aws.amazon.com/federation"
	awsTEMPConsoleUrl   = "https://console.aws.amazon.com/"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS specific commands",
	Long: `
		aws command samples:
		$tiki aws -t <ticket_path> -w   <-- Amazon console access over browser via specified ticket.
	`,
	Run: do_aws,
}

func init() {
	awsCmd.Flags().StringP("ticket", "t", "", "awsTicket to use")
	awsCmd.Flags().StringP("role", "r", "", "awsTicket role")
	awsCmd.Flags().StringP("awscli", "a", "", "path of aws client")
	awsCmd.Flags().BoolP("web", "w", false, "Amazon console access over browser")
	rootCmd.AddCommand(awsCmd)
}

func do_aws(cmd *cobra.Command, args []string) {
	profile := rootCmd.Flag("profile").Value.String()
	appconfig.AppConfig.SelectedProfile = profile

	roleCmd := cmd.Flag("role").Value.String()
	if roleCmd != "" {
		switch roleCmd {
		case "list":
			do_aws_role_list(args)
			return
		case "add":
			fmt.Printf("aws role add command is not implemented yet.\n")
			return
		case "delete":
			fmt.Printf("aws role delete command is not implemented yet.\n")
			return
		case "view":
			fmt.Printf("aws role view command is not implemented yet.\n")
			return
		default:
			fmt.Printf("Unknown aws role command: %s\n", roleCmd)
			return
		}
	}

	ticketPath, err := cmd.Flags().GetString("ticket")
	if err != nil {
		panic("Error while initializing command, Check argument(s)")
	}
	if ticketPath == "" {
		fmt.Printf("aws command requires an awsTicket provided with -t argument\n")
		return
	}

	awscli, err := cmd.Flags().GetString("awscli")
	if err != nil {
		panic("Error while initializing command, Check argument(s)")
	}

	if awscli == "" {
		awscli = appconfig.AppConfig.ViperConf.GetString(fmt.Sprintf("%s.awscli", profile))
	}

	if awscli == "" {
		fmt.Printf("aws command requires an awscli provided with -a argument\n")
		return
	}

	web, err := cmd.Flags().GetBool("web")
	if err != nil {
		panic("Error while initializing command, Check argument(s)")
	}

	//
	result, err := service.ObtainTicket(ticketPath)
	if err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while obtaining ticket: %+v\n", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	awsCreds := &models.AwsConsoleCredentials{
		SessionId:    result["AccessKeyId"].(string),
		SessionKey:   result["SecretAccessKey"].(string),
		SessionToken: result["SessionToken"].(string),
		Region:       result["Region"].(string),
	}

	var debug map[string]any

	if viper.GetBool("debug") {
		debug = map[string]interface{}{
			"Args": args,
		}
	}

	if web {
		do_web(awsCreds, debug)
	} else {
		if len(args) == 0 {
			fmt.Printf("No command provided, quitting\n")
			return
		}
		do_awscli(awsCreds, awscli, args)
	}
}

func do_aws_role_list(args []string) {
	data := service.ListAWSRoles()
	fmt.Printf("Data: %+v\n", data)
}

func do_awscli(awsCreds *models.AwsConsoleCredentials, awscli string, args []string) {
	cmd := exec.Command(awscli, args...)

	env := os.Environ()
	env = append(env, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsCreds.SessionId))
	env = append(env, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsCreds.SessionKey))
	env = append(env, fmt.Sprintf("AWS_SESSION_TOKEN=%s", awsCreds.SessionToken))
	env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", awsCreds.Region))
	cmd.Env = env

	var (
		out    bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if viper.GetBool("debug") {
		fmt.Printf("Executing aws command: [%s]\n", cmd.String())
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s", stderr.String())
	} else {
		fmt.Printf("%s", out.String())
	}
}

func do_web(
	awsCreds *models.AwsConsoleCredentials,
	debug map[string]interface{}) {

	awsCredsJSON, err := json.Marshal(awsCreds)
	if err != nil {
		errMessage := map[string]interface{}{
			"message": "marshall error",
			"details": err,
			"status":  "error",
			"debug":   debug,
		}
		utils.TikiToolResponseAsIsToJSON(errMessage)
		return
	}

	consoleUrl := fmt.Sprintf(
		"%s?Action=getSigninToken&SessionDuration=%d&Session=%s",
		awsFederationURL,
		int64(consoleSessDuration.Seconds()),
		url.QueryEscape(string(awsCredsJSON)))

	resp, err := http.Get(consoleUrl)
	if err != nil {
		fmt.Printf("AWS Federated response: %v", err)
		errMessage := map[string]interface{}{
			"details": err,
			"status":  "error",
			"message": "AWS Federated response:",
			"debug":   debug,
		}
		utils.TikiToolResponseAsIsToJSON(errMessage)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("AWS Federated response: %v", err)
		return
	}

	var federationResponse models.AwsFederationResponse
	err = json.Unmarshal(body, &federationResponse)
	if err != nil {
		fmt.Printf("Decode AWS Federated response: %v", err)
		return
	}
	loginURL := fmt.Sprintf(
		"%s?Action=login&Issuer=ticket-server&Destination=%s&SigninToken=%s",
		awsFederationURL,
		url.QueryEscape(string(awsTEMPConsoleUrl+"?region="+awsCreds.Region)),
		federationResponse.SigninToken)

	if viper.GetBool("debug") {
		fmt.Printf("Console URL opening %q...\n", loginURL) // nolint: errcheck
	}

	browser := viper.GetString("browser.preferredBrowser")
	if !utils.OpenBrowser(browser, loginURL) {
		fmt.Println("Can't open browser automatically. url:", loginURL)
	}
}
