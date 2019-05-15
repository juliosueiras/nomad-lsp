package nomadstructs

import (
	"fmt"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/juliosueiras/nomad-lsp/helper"
	"github.com/sourcegraph/go-lsp"
)

func GetAttributeCompletion(hclBlocks []*hcl.Block, result []lsp.CompletionItem) []lsp.CompletionItem {
	specType := hcldec.ImpliedType(NomadSpec)

	helper.DumpLog(len(hclBlocks))
	helper.DumpLog(hclBlocks)
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
		for _, v := range hclBlocks[1:] {
			if et := currentType.MapElementType(); et != nil {
				currentType = *et
			}

			if currentType.HasAttribute(v.Type) {
				currentType = currentType.AttributeType(v.Type)
			}
		}

		if et := currentType.MapElementType(); et != nil {
			currentType = *et
		}
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

	return result
}
