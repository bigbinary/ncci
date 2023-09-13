package deployment_targets

import (
	"fmt"

	"github.com/bigbinary/neeto-ci-cli/api/client"
	"github.com/bigbinary/neeto-ci-cli/cmd/utils"
)

func Stop(targetId string) {
	client := client.NewDeploymentTargetsV1AlphaApi()
	successful, err := client.Deactivate(targetId)
	utils.Check(err)
	if successful {
		fmt.Printf("Target [%s] was stopped successfully\n", targetId)
	}
}
