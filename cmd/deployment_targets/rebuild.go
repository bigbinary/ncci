package deployment_targets

import (
	"fmt"

	"github.com/bigbinary/neeto-ci-cli/api/client"
	"github.com/bigbinary/neeto-ci-cli/cmd/utils"
)

func Rebuild(targetId string) {
	client := client.NewDeploymentTargetsV1AlphaApi()
	successful, err := client.Activate(targetId)
	utils.Check(err)
	if successful {
		fmt.Printf("Target [%s] was rebuilt successfully\n", targetId)
	}
}
