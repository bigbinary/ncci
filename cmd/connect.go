package cmd

import (
	"fmt"
	"strings"

	"github.com/bigbinary/neeto-ci-cli/cmd/utils"
	"github.com/bigbinary/neeto-ci-cli/config"
	"github.com/spf13/cobra"

	client "github.com/bigbinary/neeto-ci-cli/api/client"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a neetoCI endpoint",
	Args:  cobra.ExactArgs(2),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		token := args[1]

		baseClient := client.NewBaseClient(token, host, "")
		client := client.NewProjectV1AlphaApiWithCustomClient(baseClient)

		_, err := client.ListProjects()

		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "%s", err)
			utils.Exit(1)
		}

		name := strings.Replace(host, ".", "_", -1)

		config.SetActiveContext(name)
		config.SetAuth(token)
		config.SetHost(host)

		cmd.Printf("connected to %s\n", host)
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
}
