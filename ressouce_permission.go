package main

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Permission dfdf
type Permission struct {
	Repo         PermissionEntry `json:"repo,omitempty"`
	Build        PermissionEntry `json:"build,omitempty"`
	RelaseBundle PermissionEntry `json:"releaseBundle,omitempty"`
}

//PermissionEntry dfdf
type PermissionEntry struct {
	IncludePatterns []string    `json:"include-patterns,omitempty"`
	ExcludePatterns []string    `json:"exclude-patterns,omitempty"`
	Repositories    []string    `json:"repositories,omitempty"`
	Actions         ActionEntry `json:"action,omitempty"`
}

//Action dfdf
type ActionEntry struct {
	Users  map[string][]string `json:"users,omitempty"`
	groups map[string][]string `json:"groups,omitempty"`
}

func resourcePermission() *schema.Resource {
	log.Printf("LAJ No Server found:")

	return &schema.Resource{
		Create: resourcePermissionCreate,
		Update: resourcePermissionUpdate,
		Read:   resourcePermissionRead,
		Delete: resourcePermissionDelete,

		Schema: map[string]*schema.Schema{
			"repo": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repositories": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"include_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"exclude_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"actions": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"groups": {
										Type:     schema.TypeString,
										Required: true,
									},
									"users": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func newPermissionFromResource(d *schema.ResourceData) *Permission {
	log.Printf("LAJ newPermissionFromResource")
	log.Printf("LAJ " + d.Get("repo.repositories").string)

	//Build:        newPermissionEntryFromResource(d, "Build"),
	//RelaseBundle: newPermissionEntryFromResource(d, "ReleaseBundle"),

	return nil
}

func resourcePermissionCreate(d *schema.ResourceData, m interface{}) error {
	newPermissionFromResource(d)
	return nil
}

func resourcePermissionRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePermissionDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePermissionUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
