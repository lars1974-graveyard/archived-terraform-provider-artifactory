package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Repository dfdf
type Repository struct {
	Key         string `json:"key,omitempty"`
	RClass      string `json:"rclass,omitempty"`
	Packagetype string `json:"packageType,omitempty"`
	Description string `json:"description,omitempty"`
}

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceRepositoryCreate,
		Update: resourceRepositoryUpdate,
		Read:   resourceRepositoryRead,
		Delete: resourceRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rclass": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "local",
			},
			"package_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "generic",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func newRepositoryFromResource(d *schema.ResourceData) *Repository {
	repository := &Repository{
		Key:         d.Get("key").(string),
		RClass:      d.Get("rclass").(string),
		Packagetype: d.Get("package_type").(string),
		Description: d.Get("description").(string),
	}
	return repository
}

func resourceRepositoryCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	repository := newRepositoryFromResource(d)
	bytedata, err := json.Marshal(repository)
	if err != nil {
		return err
	}
	key := d.Get("key").(string)
	_, err = client.Put(fmt.Sprintf("/artifactory/api/repositories/%s", key), bytes.NewBuffer(bytedata))
	if err != nil {
		return err
	}
	d.SetId(key)
	return resourceRepositoryRead(d, m)
}

func resourceRepositoryRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	projectReq, _ := client.Get(fmt.Sprintf("/artifactory/api/repositories/%s", d.Id()))
	if projectReq.StatusCode == 200 {
		var repository Repository
		body, readerr := ioutil.ReadAll(projectReq.Body)
		if readerr != nil {
			return readerr
		}
		decodeerr := json.Unmarshal(body, &repository)
		if decodeerr != nil {
			return decodeerr
		}
		d.Set("key", repository.Key)
		d.Set("rclass", repository.RClass)
		d.Set("repositorytype", repository.Packagetype)
		d.Set("description", repository.Description)
	} else {
		d.SetId("")
	}
	return nil
}

func resourceRepositoryUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	repository := newRepositoryFromResource(d)
	bytedata, err := json.Marshal(repository)
	if err != nil {
		return err
	}
	key := d.Get("key").(string)
	_, err = client.Post(fmt.Sprintf("/artifactory/api/repositories/%s", key), bytes.NewBuffer(bytedata))
	if err != nil {
		return err
	}
	d.SetId(key)
	return resourceRepositoryRead(d, m)
}

func resourceRepositoryDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	_, err := client.Delete(fmt.Sprintf("/artifactory/api/repositories/%s", d.Id()))
	if err != nil {
		return err
	}
	return nil
}
