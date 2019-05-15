package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"bitbucket.org/creachadair/jrpc2"
	"bitbucket.org/creachadair/jrpc2/channel"
	"bitbucket.org/creachadair/jrpc2/handler"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/juliosueiras/nomad-lsp/helper"
	"github.com/juliosueiras/nomad-lsp/nomadstructs"
	"github.com/sourcegraph/go-lsp"
	//"github.com/zclconf/go-cty/cty"
)

var tempFile *os.File

var Server *jrpc2.Server

var DiagsFiles = make(map[string][]lsp.Diagnostic)

func Initialize(ctx context.Context, vs lsp.InitializeParams) (lsp.InitializeResult, error) {

	file, err := ioutil.TempFile("", "tf-lsp-")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.Remove(file.Name())
	tempFile = file

	return lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			TextDocumentSync: &lsp.TextDocumentSyncOptionsOrKind{
				Options: &lsp.TextDocumentSyncOptions{
					OpenClose: true,
					Change:    1,
				},
			},
			CompletionProvider: &lsp.CompletionOptions{
				ResolveProvider:   false,
				TriggerCharacters: []string{"."},
			},
			HoverProvider: true,
			//			DocumentSymbolProvider:    true,
			//ReferencesProvider: true,
			//			DefinitionProvider:        true,
			//			DocumentHighlightProvider: true,
			//			CodeActionProvider:        true,
			//			RenameProvider:            true,
		},
	}, nil
}

func TextDocumentComplete(ctx context.Context, vs lsp.CompletionParams) (lsp.CompletionList, error) {
	var result []lsp.CompletionItem
	fileText, _ := ioutil.ReadFile(tempFile.Name())

	pos := hcl.Pos{
		Byte: helper.FindOffset(string(fileText), vs.Position.Line, vs.Position.Character),
	}

	hclBody, _ := nomadstructs.LoadHCLFile(tempFile.Name())
	blocks := hclBody.(*hclsyntax.Body).BlocksAtPos(pos)

	if blocks == nil {
		result = append(result, lsp.CompletionItem{
			Label:  "job",
			Detail: "`job` stanza",
		})
	} else {
		result = nomadstructs.GetAttributeCompletion(blocks, result)
	}

	return lsp.CompletionList{
		IsIncomplete: false,
		Items:        result,
	}, nil
}

func TextDocumentDidChange(ctx context.Context, vs lsp.DidChangeTextDocumentParams) error {
	tempFile.Truncate(0)
	tempFile.Seek(0, 0)
	tempFile.Write([]byte(vs.ContentChanges[0].Text))

	fileURL := strings.Replace(string(vs.TextDocument.URI), "file://", "", 1)
	DiagsFiles[fileURL] = nomadstructs.GetDiagnostics(tempFile.Name(), fileURL)
	TextDocumentPublishDiagnostics(Server, ctx, lsp.PublishDiagnosticsParams{
		URI:         vs.TextDocument.URI,
		Diagnostics: DiagsFiles[fileURL],
	})
	return nil
}

func TextDocumentDidOpen(ctx context.Context, vs lsp.DidOpenTextDocumentParams) error {
	fileURL := strings.Replace(string(vs.TextDocument.URI), "file://", "", 1)
	DiagsFiles[fileURL] = nomadstructs.GetDiagnostics(fileURL, fileURL)

	TextDocumentPublishDiagnostics(Server, ctx, lsp.PublishDiagnosticsParams{
		URI:         vs.TextDocument.URI,
		Diagnostics: DiagsFiles[fileURL],
	})
	tempFile.Write([]byte(vs.TextDocument.Text))
	return nil
}

func Exit(ctx context.Context, vs lsp.None) error {
	os.Remove(tempFile.Name())
	return nil
}

func TextDocumentDidClose(ctx context.Context, vs lsp.DidCloseTextDocumentParams) error {
	return nil
}

func CancelRequest(ctx context.Context, vs lsp.CancelParams) error {
	return nil
}

func TextDocumentHover(ctx context.Context, vs lsp.TextDocumentPositionParams) (lsp.Hover, error) {
	return lsp.Hover{
		Contents: []lsp.MarkedString{},
	}, nil
}

func TextDocumentPublishDiagnostics(server *jrpc2.Server, ctx context.Context, vs lsp.PublishDiagnosticsParams) error {

	return server.Push(ctx, "textDocument/publishDiagnostics", vs)
}

func main() {
	Server = jrpc2.NewServer(handler.Map{
		"initialize":              handler.New(Initialize),
		"textDocument/completion": handler.New(TextDocumentComplete),
		"textDocument/didChange":  handler.New(TextDocumentDidChange),
		"textDocument/didOpen":    handler.New(TextDocumentDidOpen),
		"textDocument/didClose":   handler.New(TextDocumentDidClose),
		"textDocument/hover":      handler.New(TextDocumentHover),
		//"textDocument/references": handler.New(TextDocumentReferences),
		//"textDocument/codeLens": handler.New(TextDocumentCodeLens),
		"exit":            handler.New(Exit),
		"$/cancelRequest": handler.New(CancelRequest),
	}, &jrpc2.ServerOptions{
		AllowPush: true,
	})

	f, err := os.OpenFile("nomad-lsp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()
	log.SetOutput(f)

	Server.Start(channel.Header("")(os.Stdin, os.Stdout))
	log.Print("Server started")

	if err := Server.Wait(); err != nil {
		log.Printf("Server exited: %v", err)
	}
}
