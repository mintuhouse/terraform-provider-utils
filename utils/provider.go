package utils

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
		// "utils_<ds_name>": dataSourceUtilsDsName(),
		},
		ResourcesMap: map[string]*schema.Resource{
		// "utils_<rs_name>": resourceUtilsRsName(),
		},
	}
}
