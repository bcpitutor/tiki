package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/bcpitutor/tiki/appconfig"
	"github.com/bcpitutor/tiki/models"
	"github.com/bcpitutor/tiki/service"
	"github.com/bcpitutor/tiki/utils"
	"github.com/spf13/cobra"
)

var ticketCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Request AWS Ticket",
	Long: `
		This command generates a time limited ticket on iTutor's AWS resources. 
	`,
	Run: do_ticket,
}

func init() {
	ticketCmd.Flags().StringP("dataFile", "f", "", "Json data file path for the ticket")
	ticketCmd.Flags().StringP("ticketType", "T", "", "TicketType, such as awsTicket")
	ticketCmd.Flags().StringP("ticket", "t", "", "Path for the ticket")
	ticketCmd.Flags().BoolP("env", "E", false, "Create env variables")
	rootCmd.AddCommand(ticketCmd)
}

func do_ticket(cmd *cobra.Command, args []string) {
	jf, _ := rootCmd.Flags().GetBool("json")

	appconfig.AppConfig.SelectedProfile = rootCmd.Flag("profile").Value.String()
	appconfig.AppConfig.JsonOutput = jf
	// debug := appconfig.AppConfig.ViperConf.GetStringMap(appconfig.AppConfig.SelectedProfile)["debug"]
	// dbool, err := strconv.ParseBool(debug.(string))
	// if err != nil {
	// 	appconfig.AppConfig.Debug = false
	// } else {
	// 	appconfig.AppConfig.Debug = dbool
	// }

	if len(args) == 0 {
		fmt.Printf("You need to specify a subcommand, such as create, info, obtain, delete or secret\n")
		return
	}

	switch args[0] {
	case "list":
		do_ticket_list(cmd, args)
	case "info":
		do_ticket_info(cmd, args)
	case "create":
		do_ticket_create(cmd)
	case "delete":
		do_ticket_delete(cmd)
	case "secret":
		do_ticket_secret(cmd, args)
	case "obtain":
		do_ticket_obtain(cmd, args)
	default:
		fmt.Printf("Unknown ticket subcommand: %s\n", args[0])
	}

}

// status: fixed
func do_ticket_list(cmd *cobra.Command, args []string) {
	result, err := service.GetTicketList()
	if err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while getting ticket list: %+v\n", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true)
	}
	newToken := result["newToken"]
	utils.CheckRefreshToken(newToken)

	intfcData, err := json.Marshal(result)
	if err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while marshalling data: %+v\n", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	var tlr models.TicketQueryResponse
	if err := json.Unmarshal(intfcData, &tlr); err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while unmarshalling data: %+v\n", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	if tlr.Status == "error" {
		utils.ErrOutput(
			tlr.Message,
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	if appconfig.AppConfig.JsonOutput {
		utils.TikiToolResponseAsIsToJSON(result)
		return
	}

	// sort the tickets by ticketPath
	sort.Slice(tlr.Tickets, func(i, j int) bool {
		return tlr.Tickets[i].TicketPath < tlr.Tickets[j].TicketPath
	})

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "TicketPath"},
			{Align: simpletable.AlignCenter, Text: "TicketType"},
			{Align: simpletable.AlignCenter, Text: "TicketRegion"},
		},
	}

	for _, row := range tlr.Tickets {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: row.TicketPath},
			{Align: simpletable.AlignLeft, Text: row.TicketType},
			{Align: simpletable.AlignCenter, Text: row.TicketRegion},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

func do_ticket_obtain(cmd *cobra.Command, args []string) {
	ticketPath, _ := cmd.Flags().GetString("ticket")
	if ticketPath == "" {
		utils.ErrOutput("You need to specify a ticket path", appconfig.AppConfig.JsonOutput, true)
	}

	if !strings.Contains(ticketPath, "/") {
		utils.ErrOutput(
			fmt.Sprintf("Bad ticket path: %s", ticketPath),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	envMode, _ := cmd.Flags().GetBool("env")

	result, err := service.ObtainTicket(ticketPath)
	if err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while obtaining ticket: %+v", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	if appconfig.AppConfig.JsonOutput {
		odata, _ := json.Marshal(result)
		fmt.Printf("%s", string(odata))
	} else {
		if envMode {
			fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", result["AccessKeyId"])
			fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", result["SecretAccessKey"])
			fmt.Printf("export AWS_SESSION_TOKEN=%s\n", result["SessionToken"])
			fmt.Printf("export AWS_REGION=%s\n", result["Region"])
			fmt.Printf("\n\n")
		} else {
			fmt.Printf("AccessKeyId: [%s]\n", result["AccessKeyId"])
			fmt.Printf("SecretAccessKey: [%s]\n", result["SecretAccessKey"])
			fmt.Printf("SessionToken: [%s]\n", result["SessionToken"])
			fmt.Printf("Region: [%s]\n", result["Region"])
		}
	}
}

func do_ticket_info(cmd *cobra.Command, args []string) {
	ticketPath, _ := cmd.Flags().GetString("ticket")
	if ticketPath == "" {
		utils.ErrOutput("You need to specify a ticket path", appconfig.AppConfig.JsonOutput, true)
	}

	if !strings.Contains(ticketPath, "/") {
		utils.ErrOutput(
			fmt.Sprintf("Bad ticket path: %s", ticketPath),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}
	result, err := service.GetTicketByPath(ticketPath)
	if err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while looking up for ticket: %+v", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	newToken := result["newToken"]
	utils.CheckRefreshToken(newToken)

	if result["status"] == "error" {
		utils.ErrOutput(
			fmt.Sprintf("Error while obtaining ticket: %+v", result["status"]),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	intfcData, err := json.Marshal(result)
	if err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while marshalling data: %+v\n", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	var tqr models.TicketQueryResponse
	if err := json.Unmarshal(intfcData, &tqr); err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while unmarshalling data: %+v\n", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	jsonFlag, _ := cmd.Flags().GetBool("json")
	if jsonFlag {
		utils.TikiToolResponseAsIsToJSON(result)
		return
	}

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Variable"},
			{Align: simpletable.AlignCenter, Text: "Value"},
		},
	}

	r := []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "TicketPath"},
		{Align: simpletable.AlignLeft, Text: tqr.Data.TicketPath},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "TicketType"},
		{Align: simpletable.AlignLeft, Text: tqr.Data.TicketType},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "TicketRegion"},
		{Align: simpletable.AlignLeft, Text: tqr.Data.TicketRegion},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	// awsTicket Specific Variables
	if tqr.Data.TicketType == "awsTicket" {
		r = []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "AssumeRole ARN"},
			{Align: simpletable.AlignLeft, Text: tqr.Data.ATSD.RoleArn},
		}
		table.Body.Cells = append(table.Body.Cells, r)

		r = []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "AssumeRole TTL"},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", tqr.Data.ATSD.TTL)},
		}
		table.Body.Cells = append(table.Body.Cells, r)

		r = []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "AWS Permissions Resource"},
			{Align: simpletable.AlignLeft, Text: tqr.Data.ATAP.Resource},
		}
		table.Body.Cells = append(table.Body.Cells, r)

		r = []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "AWS Permissions Effect"},
			{Align: simpletable.AlignLeft, Text: tqr.Data.ATAP.Effect},
		}
		table.Body.Cells = append(table.Body.Cells, r)

		for i := 0; i < len(tqr.Data.ATAP.Action); i++ {
			if i == 0 {
				r = []*simpletable.Cell{
					{Align: simpletable.AlignLeft, Text: "AWS Permissions Action"},
					{Align: simpletable.AlignLeft, Text: tqr.Data.ATAP.Action[i]},
				}
				table.Body.Cells = append(table.Body.Cells, r)
			} else {
				r = []*simpletable.Cell{
					{Align: simpletable.AlignLeft, Text: ""},
					{Align: simpletable.AlignLeft, Text: tqr.Data.ATAP.Action[i]},
				}
				table.Body.Cells = append(table.Body.Cells, r)
			}
		}

	}

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "CreatedBy"},
		{Align: simpletable.AlignLeft, Text: tqr.Data.CreatedBy},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	createDate, _ := strconv.ParseInt(tqr.Data.CreatedAt, 10, 64)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "CreatedAt"},
		{Align: simpletable.AlignLeft, Text: time.Unix(createDate, 0).String()},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "UpdatedBy"},
		{Align: simpletable.AlignLeft, Text: tqr.Data.UpdatedBy},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	updateDate, _ := strconv.ParseInt(tqr.Data.CreatedAt, 10, 64)
	r = []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "UpdatedAt"},
		{Align: simpletable.AlignLeft, Text: time.Unix(updateDate, 0).String()},
	}
	table.Body.Cells = append(table.Body.Cells, r)

	for i := 0; i < len(tqr.Data.OwnerGroups); i++ {
		if i == 0 {
			r = []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: "OwnerGroups"},
				{Align: simpletable.AlignLeft, Text: tqr.Data.OwnerGroups[i]},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		} else {
			r = []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: ""},
				{Align: simpletable.AlignLeft, Text: tqr.Data.OwnerGroups[i]},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}

func do_ticket_secret(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		utils.ErrOutput(
			"ticket secret requires a subcommand, which is either set or get",
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	ticketPath, _ := cmd.Flags().GetString("ticket")
	if ticketPath == "" {
		utils.ErrOutput("You need to specify a ticket path", appconfig.AppConfig.JsonOutput, true)
	}

	if !strings.Contains(ticketPath, "/") {
		utils.ErrOutput(
			fmt.Sprintf("Bad ticket path: %s", ticketPath),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	switch args[1] {
	case "set":
		if len(args) == 2 {
			utils.ErrOutput(
				"ticket secret set subcommand requires a value",
				appconfig.AppConfig.JsonOutput,
				true,
			)
		}
		result, err := service.SetTicketSecret(ticketPath, args[2])
		if err != nil {
			utils.ErrOutput(
				fmt.Sprintf("Error while setting ticket secret: %+v", err.Error()),
				appconfig.AppConfig.JsonOutput,
				true,
			)
		}

		newToken := result["newToken"]
		utils.CheckRefreshToken(newToken)

		fmt.Printf("%+s\n", result["status"])

	case "get":
		result, err := service.GetTicketSecret(ticketPath)
		if err != nil {
			utils.ErrOutput(
				fmt.Sprintf("Error while getting ticket secret: %+v", err.Error()),
				appconfig.AppConfig.JsonOutput,
				true,
			)
		}

		newToken := result["newToken"]
		utils.CheckRefreshToken(newToken)

		// TODO: check if secret ends with a new line char
		if appconfig.AppConfig.JsonOutput {
			type TicketSecretOutput struct {
				Secret string `json:"secret"`
			}

			tso := TicketSecretOutput{
				Secret: result["message"].(string),
			}

			odata, _ := json.Marshal(tso)
			fmt.Printf("%s", string(odata))
		} else {
			fmt.Printf("%s", result["message"])
		}

	default:
		utils.ErrOutput(
			fmt.Sprintf("Unknown ticket secret subcommand %s\n", args[1]),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}
}

func do_ticket_create(cmd *cobra.Command) {
	ticketDataJsonFilePath, err := cmd.Flags().GetString("dataFile")
	if err != nil || ticketDataJsonFilePath == "" {
		fmt.Printf("You need to specify a dataFile with -f switch, in non-interactive mode\n")
		return
	}

	result, err := service.CreateTicket(ticketDataJsonFilePath)
	utils.ErrOutput(
		fmt.Sprintf("Error while creating ticket: %+v\n", err.Error()),
		appconfig.AppConfig.JsonOutput,
		true,
	)

	newToken := result["newToken"]
	utils.CheckRefreshToken(newToken)

	if result["status"] != "success" {
		fmt.Printf("%+s\n", result["message"])
	} else {
		fmt.Printf("Ticket created successfully\n")
	}
}

func do_ticket_delete(cmd *cobra.Command) {

	ticketPath, err := cmd.Flags().GetString("ticket")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	if ticketPath == "" {
		fmt.Printf("You need to specify a ticket path\n")
		return
	}

	if !strings.Contains(ticketPath, "/") {
		fmt.Printf("Bad ticket path: %s\n", ticketPath)
		return
	}

	result := service.DeleteTicket(ticketPath)
	newToken := result["newToken"]
	utils.CheckRefreshToken(newToken)

	fmt.Printf("Ticket %s deleted\n", ticketPath)
}
