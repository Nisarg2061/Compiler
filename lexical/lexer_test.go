package main

import (
	"reflect"
	"testing"
)

func TestLexerKeywords(t *testing.T) {
	input := "if else "
	lexer := NewLexer(input)

	expectedTokens := []Token{
		{Type: TokenKeyword, Value: "if"},
		{Type: TokenKeyword, Value: "else"},
		{Type: TokenEOF, Value: ""},
	}

	var actualTokens []Token
	for {
		token := lexer.GetNextToken()
		actualTokens = append(actualTokens, token)
		if token.Type == TokenEOF {
			break
		}
	}

	if !reflect.DeepEqual(expectedTokens, actualTokens) {
		t.Errorf("Expected tokens %v, got %v", expectedTokens, actualTokens)
	}
}

func TestLexerIdentifiersAndNumbers(t *testing.T) {
	input := "var1 123 "
	lexer := NewLexer(input)

	expectedTokens := []Token{
		{Type: TokenIdentifier, Value: "var1"},
		{Type: TokenNumber, Value: "123"},
		{Type: TokenEOF, Value: ""},
	}

	var actualTokens []Token
	for {
		token := lexer.GetNextToken()
		actualTokens = append(actualTokens, token)
		if token.Type == TokenEOF {
			break
		}
	}

	if !reflect.DeepEqual(expectedTokens, actualTokens) {
		t.Errorf("Expected tokens %v, got %v", expectedTokens, actualTokens)
	}
}

func TestLexerOperators(t *testing.T) {
	input := "+ * - /"
	lexer := NewLexer(input)

	expectedTokens := []Token{
		{Type: TokenOperator, Value: "+"},
		{Type: TokenOperator, Value: "*"},
		{Type: TokenOperator, Value: "-"},
		{Type: TokenOperator, Value: "/"},
		{Type: TokenEOF, Value: ""},
	}

	var actualTokens []Token
	for {
		token := lexer.GetNextToken()
		actualTokens = append(actualTokens, token)
		if token.Type == TokenEOF {
			break
		}
	}

	if !reflect.DeepEqual(expectedTokens, actualTokens) {
		t.Errorf("Expected tokens %v, got %v", expectedTokens, actualTokens)
	}
}
