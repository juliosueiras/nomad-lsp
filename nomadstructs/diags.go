package nomadstructs

import (
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/juliosueiras/nomad-lsp/helper"
	"github.com/sourcegraph/go-lsp"
	"os"
)

func GetDiagnostics(fileName string, originalFile string) []lsp.Diagnostic {
	result := make([]lsp.Diagnostic, 0)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return result
	}

	if _, err := os.Stat(originalFile); os.IsNotExist(err) {
		originalFile = fileName
	}

	hclBody, hclDiags := LoadHCLFile(fileName)
	helper.DumpLog(hclDiags)

	for _, diag := range hclDiags {
		result = append(result, lsp.Diagnostic{
			Severity: lsp.DiagnosticSeverity(diag.Severity),
			Message:  diag.Detail,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      diag.Subject.Start.Line - 1,
					Character: diag.Subject.Start.Column - 1,
				},
				End: lsp.Position{
					Line:      diag.Subject.End.Line - 1,
					Character: diag.Subject.End.Column - 1,
				},
			},
			Source: "HCL",
		})
	}

	_, nomadDiags := hcldec.Decode(hclBody, NomadSpec, &hcl.EvalContext{})

	for _, diag := range nomadDiags {
		result = append(result, lsp.Diagnostic{
			Severity: lsp.DiagnosticSeverity(diag.Severity),
			Message:  diag.Detail,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      diag.Subject.Start.Line - 1,
					Character: diag.Subject.Start.Column - 1,
				},
				End: lsp.Position{
					Line:      diag.Subject.End.Line - 1,
					Character: diag.Subject.End.Column - 1,
				},
			},
			Source: "Nomad",
		})
	}

	return result
}
