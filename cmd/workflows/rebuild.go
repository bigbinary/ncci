package workflows

import (
	"fmt"

	"github.com/bigbinary/ncci/api/client"
	"github.com/bigbinary/ncci/cmd/utils"
)

func Rebuild(id string) {
	wfClient := client.NewWorkflowV1AlphaApi()
	body, err := wfClient.Rebuild(id)
	utils.Check(err)
	fmt.Printf("%s\n", string(body))
}
