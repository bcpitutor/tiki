package cmd

import (
	"encoding/json"
	"fmt"
	"os"
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
	rootCmd.AddCommand(ticketCmd)
}

func do_ticket(cmd *cobra.Command, args []string) {
	profile := rootCmd.Flag("profile").Value.String()
	jf, _ := rootCmd.Flags().GetBool("json")

	appconfig.AppConfig.SelectedProfile = profile
	appconfig.AppConfig.JsonOutput = jf
	debug := appconfig.AppConfig.ViperConf.GetStringMap(profile)["debug"]
	dbool, err := strconv.ParseBool(debug.(string))
	if err != nil {
		appconfig.AppConfig.Debug = false
	} else {
		appconfig.AppConfig.Debug = dbool
	}

	if len(args) == 0 {
		fmt.Printf("You need to specify a subcommand, such as create, obtain, delete or secret\n")
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
			fmt.Sprintf("Bad ticket path: %s\n", ticketPath),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	result, err := service.ObtainTicket(ticketPath)
	if err != nil {
		utils.ErrOutput(
			fmt.Sprintf("Error while obtaining ticket: %+v\n", err.Error()),
			appconfig.AppConfig.JsonOutput,
			true,
		)
	}

	if appconfig.AppConfig.JsonOutput {
		odata, _ := json.Marshal(result)
		fmt.Printf("%s", string(odata))
	} else {
		fmt.Printf("Hayde: Hoppa\n")
		fmt.Printf("AccessKeyId: [%s]\n", result["AccessKeyId"])
		fmt.Printf("SecretAccessKey: [%s]\n", result["SecretAccessKey"])
		fmt.Printf("SessionToken: [%s]\n", result["SessionToken"])
		fmt.Printf("Region: [%s]\n", result["Region"])
	}
}

func do_ticket_secret(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		fmt.Printf("ticket secret requires a subcommand, which is either set or get\n")
		return
	}
	ticketPath, _ := cmd.Flags().GetString("ticket")

	switch args[1] {
	case "set":
		if len(args) == 2 {
			fmt.Printf("ticket secret set subcommand requires a value\n")
			return
		}
		result := service.SetTicketSecret(ticketPath, args[2])
		newToken := result["newToken"]
		utils.CheckRefreshToken(newToken)

		fmt.Printf("%+s\n", result["status"])

	case "get":
		result, err := service.GetTicketSecret(ticketPath)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(-1)
		}

		newToken := result["newToken"]
		utils.CheckRefreshToken(newToken)

		fmt.Println(result["message"])
	default:
		fmt.Printf("Unknown ticket secret subcommand %s\n", args[1])
		return
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

func do_ticket_info(cmd *cobra.Command, args []string) {
	ticketPath, _ := cmd.Flags().GetString("ticket")
	if !strings.Contains(ticketPath, "/") {
		fmt.Printf("Bad ticket path: %s\n", ticketPath)
		return
	}
	result := service.GetTicketByPath(ticketPath)
	newToken := result["newToken"]
	utils.CheckRefreshToken(newToken)

	if result["status"] == "error" {
		fmt.Println(result["message"])
		return
	}

	intfcData, err := json.Marshal(result)
	if err != nil {
		fmt.Println("err1:", err)
	}
	//zlogger.Log.Info("intfcData: " + string(intfcData))

	var tqr models.TicketQueryResponse
	if err := json.Unmarshal(intfcData, &tqr); err != nil {
		panic(err)
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

func do_ticket_delete(cmd *cobra.Command) {
	ticketPath, err := cmd.Flags().GetString("ticket")
	if err != nil || ticketPath == "" {
		fmt.Printf("Error: %s\n", err)
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
