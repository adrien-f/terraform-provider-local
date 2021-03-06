package local

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"local_file":      resourceLocalFile(),
			"local_directory": resourceLocalDirectory(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"local_file":      dataSourceLocalFile(),
			"local_directory": dataSourceLocalDirectory(),
		},
	}
}
