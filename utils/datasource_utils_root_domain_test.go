package utils

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccUtilsRootDomainDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckUtilsRootDomainDataSourceConfig("com"),
				ExpectError: regexp.MustCompile(`com is a suffix`),
			},
			{
				Config: testAccCheckUtilsRootDomainDataSourceConfig("example.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.utils_root_domain.test", "name", "example.com"),
				),
			},
			{
				Config: testAccCheckUtilsRootDomainDataSourceConfig("www.example.com"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.utils_root_domain.test", "name", "example.com"),
				),
			},
		},
	})
}

func testAccCheckUtilsRootDomainDataSourceConfig(domain string) string {
	return fmt.Sprintf(`
data "utils_root_domain" "test" {
	domain = "%s"
}
`, domain)
}

func test(s *terraform.State) {
	_ = s
}
