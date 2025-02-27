// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/healthcare/Dataset.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package healthcare

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceHealthcareDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcareDatasetCreate,
		Read:   resourceHealthcareDatasetRead,
		Update: resourceHealthcareDatasetUpdate,
		Delete: resourceHealthcareDatasetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceHealthcareDatasetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the Dataset.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name for the Dataset.`,
			},
			"encryption_spec": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `A nested object resource.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `KMS encryption key that is used to secure this dataset and its sub-resources. The key used for
encryption and the dataset must be in the same location. If empty, the default Google encryption
key will be used to secure this dataset. The format is
projects/{projectId}/locations/{locationId}/keyRings/{keyRingId}/cryptoKeys/{keyId}.`,
						},
					},
				},
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
"America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
(e.g., HL7 messages) where no explicit timezone is specified.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The fully qualified name of this dataset`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceHealthcareDatasetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareDatasetName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	timeZoneProp, err := expandHealthcareDatasetTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	encryptionSpecProp, err := expandHealthcareDatasetEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionSpecProp)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets?datasetId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Dataset: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Dataset: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.HealthcareDatasetNotInitialized},
	})
	if err != nil {
		return fmt.Errorf("Error creating Dataset: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Dataset %q: %#v", d.Id(), res)

	return resourceHealthcareDatasetRead(d, meta)
}

func resourceHealthcareDatasetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Dataset: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.HealthcareDatasetNotInitialized},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("HealthcareDataset %q", d.Id()))
	}

	res, err = resourceHealthcareDatasetDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing HealthcareDataset because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	if err := d.Set("name", flattenHealthcareDatasetName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("time_zone", flattenHealthcareDatasetTimeZone(res["timeZone"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}
	if err := d.Set("encryption_spec", flattenHealthcareDatasetEncryptionSpec(res["encryptionSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Dataset: %s", err)
	}

	return nil
}

func resourceHealthcareDatasetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Dataset: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	timeZoneProp, err := expandHealthcareDatasetTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Dataset %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("time_zone") {
		updateMask = append(updateMask, "timeZone")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:               config,
			Method:               "PATCH",
			Project:              billingProject,
			RawURL:               url,
			UserAgent:            userAgent,
			Body:                 obj,
			Timeout:              d.Timeout(schema.TimeoutUpdate),
			Headers:              headers,
			ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.HealthcareDatasetNotInitialized},
		})

		if err != nil {
			return fmt.Errorf("Error updating Dataset %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Dataset %q: %#v", d.Id(), res)
		}

	}

	return resourceHealthcareDatasetRead(d, meta)
}

func resourceHealthcareDatasetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Dataset: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{HealthcareBasePath}}projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Dataset %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.HealthcareDatasetNotInitialized},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Dataset")
	}

	log.Printf("[DEBUG] Finished deleting Dataset %q: %#v", d.Id(), res)
	return nil
}

func resourceHealthcareDatasetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/datasets/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/datasets/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenHealthcareDatasetName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareDatasetTimeZone(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenHealthcareDatasetEncryptionSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenHealthcareDatasetEncryptionSpecKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenHealthcareDatasetEncryptionSpecKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandHealthcareDatasetName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareDatasetTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareDatasetEncryptionSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandHealthcareDatasetEncryptionSpecKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandHealthcareDatasetEncryptionSpecKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceHealthcareDatasetDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	if err := d.Set("self_link", res["name"].(string)); err != nil {
		return nil, fmt.Errorf("Error setting self_link: %s", err)
	}
	res["name"] = d.Get("name").(string)
	return res, nil
}
