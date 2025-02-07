package workflows

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	client "github.com/bigbinary/neeto-ci-cli/api/client"
	"github.com/bigbinary/neeto-ci-cli/api/models"
	"github.com/bigbinary/neeto-ci-cli/cmd/utils"
)

func List(projectID string) {
	wfClient := client.NewWorkflowV1AlphaApi()
	workflows, err := wfClient.ListWorkflows(projectID)
	utils.Check(err)

	prettyPrint(workflows)
}

func prettyPrint(workflows *models.WorkflowListV1Alpha) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

	fmt.Fprintln(w, "WORKFLOW ID\tINITIAL PIPELINE ID\tCREATION TIME\tLABEL")

	for _, p := range workflows.Workflow {
		createdAt := time.Unix(p.CreatedAt.Seconds, 0).Format("2006-01-02 15:04:05")

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", p.Id, p.InitialPplId, createdAt, p.BranchName)
	}

	if err := w.Flush(); err != nil {
		fmt.Printf("Error flushing when pretty printing workflows: %v\n", err)
	}
}
