package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mintuhouse/terraform-provider-utils/utils"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: utils.Provider})
}
