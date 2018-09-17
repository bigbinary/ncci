package cmd

import (
	"fmt"

	client "github.com/semaphoreci/cli/api/client"
	"github.com/semaphoreci/cli/cmd/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [KIND] [NAME]",
	Short: "Delete a resource.",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
}

var DeleteDashboardCmd = &cobra.Command{
	Use:     "dashboard [NAME]",
	Short:   "Delete a dashboard.",
	Long:    ``,
	Aliases: []string{"dashboards", "dash"},
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		c := client.NewDashboardV1AlphaApi()

		err := c.DeleteDashboard(name)

		utils.Check(err)

		fmt.Printf("Dashboard '%s' deleted.\n", name)
	},
}

var DeleteSecretCmd = &cobra.Command{
	Use:     "secret [NAME]",
	Short:   "Delete a secret.",
	Long:    ``,
	Aliases: []string{"secrets"},
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		c := client.NewSecretV1BetaApi()

		err := c.DeleteSecret(name)

		utils.Check(err)

		fmt.Printf("Secret '%s' deleted.\n", name)
	},
}

var DeleteProjectCmd = &cobra.Command{
	Use:     "project [NAME]",
	Short:   "Delete a project.",
	Long:    ``,
	Aliases: []string{"projects", "prj"},
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		c := client.NewProjectV1AlphaApi()

		err := c.DeleteProject(name)

		utils.Check(err)

		fmt.Printf("Project '%s' deleted.\n", name)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)

	deleteCmd.AddCommand(DeleteDashboardCmd)
	deleteCmd.AddCommand(DeleteProjectCmd)
	deleteCmd.AddCommand(DeleteSecretCmd)
}
