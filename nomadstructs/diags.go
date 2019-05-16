package nomadstructs

import (
	"fmt"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/sourcegraph/go-lsp"
	"os"
	"reflect"
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

	// Need check for docker mounts fix
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

	for _, v := range hclBody.(*hclsyntax.Body).Blocks {
		for _, p := range v.Body.Blocks {
			if p.Type == "group" {
				for _, m := range p.Body.Blocks {
					if m.Type == "task" {
						if driverAttr := m.Body.Attributes["driver"]; driverAttr != nil {
							if reflect.TypeOf(driverAttr.Expr) == reflect.TypeOf(&hclsyntax.TemplateExpr{}) {
								driverName := driverAttr.Expr.(*hclsyntax.TemplateExpr).Parts[0].(*hclsyntax.LiteralValueExpr).Val.AsString()
								for i, n := range m.Body.Blocks {
									if n.Type == "config" {
										if driverSpec := GetDriverSpec(driverName); driverSpec != nil {
											body := SanitizeDriverConfig(n.Body, driverName)
											_, driverDiags := hcldec.Decode(body, driverSpec, &hcl.EvalContext{})

											for _, diag := range driverDiags {
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
													Source: fmt.Sprintf(" Task Driver(%s)", driverName),
												})
											}
										}
										m.Body.Blocks = append(m.Body.Blocks[:i], m.Body.Blocks[i+1:]...)

									}
								}
							}
						}
					}
				}
			}
		}
	}

	_, nomadDiags := hcldec.Decode(hclBody, JobSpec, &hcl.EvalContext{})

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
