package main

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//This is a comment
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("ARTIFACTORY_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARTIFACTORY_PASSWORD", nil),
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"artifactory_repository": resourceRepository(),
			"artifactory_permission": resourcePermission(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := &Client{
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
		HTTPClient: &http.Client{},
	}

	return client, nil
}
