package pipelines

import (
	"fmt"

	client "github.com/bigbinary/ncci/api/client"
	"github.com/bigbinary/ncci/cmd/utils"
)

func Rebuild(id string) {
	c := client.NewPipelinesV1AlphaApi()
	body, err := c.PartialRebuildPpl(id)
	utils.Check(err)
	fmt.Printf("%s\n", string(body))
}
