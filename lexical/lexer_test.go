package main

import (
	"reflect"
	"testing"
)

// Helper function to compare actual tokens with expected ones and report errors
func compareTokens(t *testing.T, input string, expectedTokens []Token) {
	lexer := NewLexer(input)
	var actualTokens []Token

	for {
		token := lexer.GetNextToken()
		actualTokens = append(actualTokens, token)
		if token.Type == TokenEOF {
			break
		}
	}

	if !reflect.DeepEqual(expectedTokens, actualTokens) {
		t.Errorf("For input %q, expected tokens %v, but got %v", input, expectedTokens, actualTokens)
	}
}

func TestLexer(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []Token
	}{
		{
			name:  "Keywords",
			input: "if else ",
			expectedTokens: []Token{
				{Type: TokenKeyword, Value: "if"},
				{Type: TokenKeyword, Value: "else"},
				{Type: TokenEOF, Value: ""},
			},
		},
		{
			name:  "Identifiers and Numbers",
			input: "var1 123 ",
			expectedTokens: []Token{
				{Type: TokenIdentifier, Value: "var1"},
				{Type: TokenNumber, Value: "123"},
				{Type: TokenEOF, Value: ""},
			},
		},
		{
			name:  "Operators",
			input: "+ * - /",
			expectedTokens: []Token{
				{Type: TokenOperator, Value: "+"},
				{Type: TokenOperator, Value: "*"},
				{Type: TokenOperator, Value: "-"},
				{Type: TokenOperator, Value: "/"},
				{Type: TokenEOF, Value: ""},
			},
		},
		{
			name:  "Empty Input",
			input: "",
			expectedTokens: []Token{
				{Type: TokenEOF, Value: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compareTokens(t, tt.input, tt.expectedTokens)
		})
	}
}

