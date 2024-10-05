package main

import (
	"unicode"
)

type TokenType string

const (
	TokenKeyword    TokenType = "KEYWORD"
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenOperator   TokenType = "OPERATOR"
	TokenEOF        TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input       string
	position    int
	currentChar rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance() // Prime the lexer
	return l
}

func (l *Lexer) advance() {
	if l.position < len(l.input) {
		l.currentChar = rune(l.input[l.position])
		l.position++
	} else {
		l.currentChar = 0 // EOF
	}
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *Lexer) handleIdentifier() Token {
	start := l.position - 1
	for isLetter(l.currentChar) || isDigit(l.currentChar) {
		l.advance()
	}
	value := l.input[start : l.position-1]

	// Define keywords for scalability
	keywords := map[string]TokenType{
		"if":   TokenKeyword,
		"else": TokenKeyword,
	}

	// Check if the value is a keyword
	if tokenType, isKeyword := keywords[value]; isKeyword {
		return Token{Type: tokenType, Value: value}
	}

	return Token{Type: TokenIdentifier, Value: value}
}

// GetNextToken retrieves the next token from the input.
func (l *Lexer) GetNextToken() Token {
	for l.currentChar != 0 {
		if unicode.IsSpace(rune(l.currentChar)) {
			l.advance()
			continue
		}

		if isLetter(l.currentChar) {
			return l.handleIdentifier()
		}

		if isDigit(l.currentChar) {
			return l.handleNumber()
		}

		if l.currentChar == '+' || l.currentChar == '*' || l.currentChar == '-' || l.currentChar == '/' {
			op := string(l.currentChar)
			l.advance()
			return Token{Type: TokenOperator, Value: op}
		}

		l.advance() // advance to the next character
	}

	return Token{Type: TokenEOF, Value: ""}
}

// handleNumber processes numbers (integers for simplicity).
func (l *Lexer) handleNumber() Token {
	start := l.position - 1
	for isDigit(l.currentChar) {
		l.advance()
		if unicode.IsSpace(rune(l.currentChar)) {
			break
		}
	}
	l.position-- // Move back to the last digit of the number
	value := l.input[start:l.position]

	return Token{Type: TokenNumber, Value: value}
}
