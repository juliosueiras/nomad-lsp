package nomadstructs

import (
	"fmt"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/sourcegraph/go-lsp"
)

func GetAttributeCompletion(hclBlocks []*hcl.Block, result []lsp.CompletionItem) []lsp.CompletionItem {
	specType := hcldec.ImpliedType(JobSpec)

	if len(hclBlocks) == 1 {
		for k, v := range specType.ElementType().AttributeTypes() {
			if v.IsObjectType() {
				result = append(result, lsp.CompletionItem{
					Label:  k,
					Detail: " stanza",
				})
			} else {
				result = append(result, lsp.CompletionItem{
					Label:  k,
					Detail: fmt.Sprintf(" %s", v.FriendlyName()),
				})
			}
		}
	} else {
		currentType := specType
		prevBlock := hclBlocks[0]
		for _, v := range hclBlocks[1:] {
			if v.Type == "config" {
				s := prevBlock.Body.(*hclsyntax.Body)
				if s.Attributes["driver"] != nil {
					if ct := GetDriverSpec(s.Attributes["driver"].Expr.(*hclsyntax.TemplateExpr).Parts[0].(*hclsyntax.LiteralValueExpr).Val.AsString()); ct != nil {
						currentType = hcldec.ImpliedType(ct)
					}
				}
				continue
			}

			if et := currentType.MapElementType(); et != nil {
				currentType = *et
			}

			if currentType.HasAttribute(v.Type) {
				currentType = currentType.AttributeType(v.Type)
			}

			prevBlock = v
		}

		if et := currentType.MapElementType(); et != nil {
			currentType = *et
		}

		if !currentType.IsPrimitiveType() {
			for k, v := range currentType.AttributeTypes() {
				if v.IsObjectType() {
					result = append(result, lsp.CompletionItem{
						Label:  k,
						Detail: " stanza",
					})
				} else {
					result = append(result, lsp.CompletionItem{
						Label:  k,
						Detail: fmt.Sprintf(" %s", v.FriendlyName()),
					})
				}
			}
		}
	}

	return result
}
