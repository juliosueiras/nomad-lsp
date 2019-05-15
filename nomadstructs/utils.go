package nomadstructs

import (
	"fmt"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/spf13/afero"
	"strings"
)

// from https://github.com/hashicorp/terraform/configs/parser.go#L49
func LoadHCLFile(path string) (hcl.Body, hcl.Diagnostics) {
	src, err := afero.Afero{Fs: afero.OsFs{}}.ReadFile(path)
	p := hclparse.NewParser()

	if err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}

	var file *hcl.File
	var diags hcl.Diagnostics
	switch {
	case strings.HasSuffix(path, ".json"):
		file, diags = p.ParseJSON(src, path)
	default:
		file, diags = p.ParseHCL(src, path)
	}

	if file == nil || file.Body == nil {
		return hcl.EmptyBody(), diags
	}

	return file.Body, diags
}
