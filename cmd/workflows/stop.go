package workflows

import (
	"fmt"

	"github.com/bigbinary/ncci/api/client"
	"github.com/bigbinary/ncci/cmd/utils"
)

func Stop(id string) {
	c := client.NewWorkflowV1AlphaApi()
	body, err := c.StopWf(id)
	utils.Check(err)
	fmt.Printf("%s\n", string(body))
}
