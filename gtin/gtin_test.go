// Package gtin_test contains tests for the gtin grammar.
// The tests should be run with the -timeout flag, to ensure the parser doesn't
// get stuck.
//
// Do not edit this file, it is generated by maketest.go
//
package gtin_test

import (
	"bramp.net/antlr4test-go/gtin"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
	"testing"
)

const MAX_TOKENS = 1000000

var examples = []string{
	"grammars-v4/gtin/examples/bookland1.txt",
	"grammars-v4/gtin/examples/bookland1.txt.tree",
	"grammars-v4/gtin/examples/bookland2.txt",
	"grammars-v4/gtin/examples/bookland2.txt.tree",
	"grammars-v4/gtin/examples/ean8.txt",
	"grammars-v4/gtin/examples/ean8.txt.tree",
	"grammars-v4/gtin/examples/gtin13.txt",
	"grammars-v4/gtin/examples/gtin13.txt.tree",
	"grammars-v4/gtin/examples/gtin14_1.txt",
	"grammars-v4/gtin/examples/gtin14_1.txt.tree",
	"grammars-v4/gtin/examples/ismn.txt",
	"grammars-v4/gtin/examples/isnm.txt.tree",
	"grammars-v4/gtin/examples/issn.txt",
	"grammars-v4/gtin/examples/issn.txt.tree",
	"grammars-v4/gtin/examples/sup2.txt",
	"grammars-v4/gtin/examples/sup2.txt.tree",
	"grammars-v4/gtin/examples/sup5.txt",
	"grammars-v4/gtin/examples/sup5.txt.tree",
	"grammars-v4/gtin/examples/upc_a_1.txt",
	"grammars-v4/gtin/examples/upc_a_1.txt.tree",
	"grammars-v4/gtin/examples/upc_a_1_hyphen.txt",
	"grammars-v4/gtin/examples/upc_a_1_hyphen.txt.tree",
	"grammars-v4/gtin/examples/upc_e_1.txt",
	"grammars-v4/gtin/examples/upc_e_1.txt.tree",
}

func newCharStream(filename string) (antlr.CharStream, error) {
	var input antlr.CharStream
	input, err := antlr.NewFileStream(filepath.Join("..", filename))
	if err != nil {
		return nil, err
	}

	return input, nil
}

// TODO Add an Example func

func TestgtinLexer(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := gtin.NewgtinLexer(input)

		// Try and read all tokens
		i := 0
		for ; i < MAX_TOKENS; i++ {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
		}

		// If we read too many tokens, then perhaps there is a problem with the lexer.
		if i == MAX_TOKENS {
			t.Errorf("NewgtinLexer(%q) read %d tokens without finding EOF", file, i)
		}
	}
}

func TestgtinParser(t *testing.T) {
	for _, file := range examples {
		input, err := newCharStream(file)
		if err != nil {
			t.Errorf("Failed to open example file: %s", err)
		}

		// Create the Lexer
		lexer := gtin.NewgtinLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := gtin.NewgtinParser(stream)
		p.BuildParseTrees = true
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true)) // TODO Change this
		p.AddErrorListener(antlr.NewConsoleErrorListener())

		// Finally test
		p.Gtin()

		// TODO Check for errors
	}
}