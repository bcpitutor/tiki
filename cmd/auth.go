package cmd

import (
	"fmt"

	"github.com/bcpitutor/tiki/service"
	"github.com/bcpitutor/tiki/subcmd"
	"github.com/bcpitutor/tiki/utils"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication over google via tikiserver",
	Long: `
		the default port of listener is 3434. If it is a reason of conflict. you can use;
		
		$ tiki auth -P <port>
	`,
	Run: do_auth,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	subcmd.GoogleAuthenticationViaTikiServer(cmd, args)
	// },
}

func do_auth(cmd *cobra.Command, args []string) {
	profile := rootCmd.Flag("profile").Value.String()

	renew, err := cmd.Flags().GetBool("renew")
	//profile := rootCmd.Flag("profile").Value.String()
	if err != nil {
		panic("Error while initializing command, Check argument(s)")
	}

	if renew {
		do_auth_renew()
		return
	}

	subcmd.GoogleAuthenticationViaTikiServer(profile, cmd, args)
}

func do_auth_renew() {
	// todo: ??
	result := service.AuthRenew()
	newToken := result["newToken"]
	utils.CheckRefreshToken(newToken)

	fmt.Printf("Result: %+v\n", result)

}

func init() {
	authCmd.Flags().StringP("listenerPort", "P", "3434", "Local listener port.")
	authCmd.Flags().BoolP("renew", "r", false, "renew token, even though it is not expired.")
	rootCmd.AddCommand(authCmd)
}
