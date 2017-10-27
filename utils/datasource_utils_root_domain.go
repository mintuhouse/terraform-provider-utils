package utils

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/weppos/publicsuffix-go/publicsuffix"
)

func dataSourceUtilsRootDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUtilsRootDomainRead,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUtilsRootDomainRead(d *schema.ResourceData, meta interface{}) error {
	domainName := d.Get("domain").(string)
	rootDomainName, err := publicsuffix.Domain(domainName)

	if err != nil {
		return err
	}

	d.Set("name", rootDomainName)
	d.SetId(domainName)

	return nil
}
