// Package singletonlist keeps Terraform Plugin Framework single nested blocks
// as singleton lists in the generated CRDs.
package singletonlist

import (
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Keep generates the single nested blocks at the given dot-separated
// Terraform paths (e.g. "config", "config.jira") as singleton lists, the shape
// upjet generated before it mapped single nested blocks to embedded objects.
// Nested paths must be listed together with all their ancestors. At runtime
// the lists are converted to the objects the Plugin Framework resource
// expects, and back.
func Keep(r *config.Resource, paths ...string) {
	if !r.ShouldUseTerraformPluginFrameworkClient() {
		// SDKv2 resources read r.TerraformResource at runtime, so the schema
		// must not be mutated for them.
		panic(fmt.Sprintf("singletonlist: %s is not a Plugin Framework resource", r.Name))
	}
	convPaths := make([]string, 0, len(paths))
	for _, p := range paths {
		s := config.GetSchema(r.TerraformResource, p)
		if s == nil {
			panic(fmt.Sprintf("singletonlist: %s has no Terraform schema at path %q", r.Name, p))
		}
		s.Type, s.MaxItems = schema.TypeList, 1
		convPaths = append(convPaths, strings.ReplaceAll(p, ".", "[*]."))
	}
	r.TerraformConversions = append(r.TerraformConversions, listToObject(convPaths))
}

// listToObject converts between the CRD singleton lists and the Terraform
// objects at the given paths: the reverse of upjet's NewTFSingletonConversion.
type listToObject []string

func (paths listToObject) Convert(params map[string]any, _ *config.Resource, mode config.Mode) (map[string]any, error) {
	if mode == config.ToTerraform {
		return conversion.Convert(params, paths, conversion.ToEmbeddedObject, nil)
	}
	return conversion.Convert(params, paths, conversion.ToSingletonList, nil)
}
