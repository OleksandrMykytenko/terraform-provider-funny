package funny

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		// provider parameters here, like AWS credentials or region for AWS provider
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "token should be provided here",
			},
		},
		// map with resource names and their constructors
		ResourcesMap: map[string]*schema.Resource{
			"funny_object": resourceObject(),
		},
	}
}
