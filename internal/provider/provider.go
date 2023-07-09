package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider exported function.
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
	}
}
