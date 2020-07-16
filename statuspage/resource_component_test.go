package statuspage

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccStatuspageComponentBasic(t *testing.T) {
	rid := acctest.RandIntRange(1000, 9999)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComponentBasic(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "test component"),
					resource.TestCheckResourceAttr("statuspage_component.default", "status", "operational"),
					resource.TestCheckResourceAttr("statuspage_component.default", "showcase", "true"),
				),
			},
			{
				Config: testAccComponentUpdate(rid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("statuspage_component.default", "id"),
					resource.TestCheckResourceAttr("statuspage_component.default", "description", "updated component"),
					resource.TestCheckResourceAttr("statuspage_component.default", "status", "major_outage"),
					resource.TestCheckResourceAttr("statuspage_component.default", "showcase", "false"),
				),
			},
		},
	})
}

func testAccComponentBasic(rand int) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "tf-testacc-component-%d"
	}

	variable "pageid" {
		default = "%s"
	}

	resource "statuspage_component" "default" {
		page_id = "${var.pageid}"
		name = "${var.name}"
		description = "test component"
		status = "operational"
		showcase = true
	}
	`, rand, pageID)
}

func testAccComponentUpdate(rand int) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "tf-testacc-component-%d"
	}

	variable "pageid" {
		default = "%s"
	}

	resource "statuspage_component" "default" {
		page_id = "${var.pageid}"
		name = "${var.name}"
		description = "updated component"
		status = "major_outage"
		showcase = false
	}
	`, rand, pageID)
}
