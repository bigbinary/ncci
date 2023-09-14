package workflows

import (
	"fmt"

	"github.com/bigbinary/neeto-ci-cli/api/client"
	"github.com/bigbinary/neeto-ci-cli/cmd/utils"
)

func Stop(id string) {
	c := client.NewWorkflowV1AlphaApi()
	body, err := c.StopWf(id)
	utils.Check(err)
	fmt.Printf("%s\n", string(body))
}
