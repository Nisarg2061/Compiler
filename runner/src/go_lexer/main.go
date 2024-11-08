package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type TokenType string

const (
	TokenKeyword    TokenType = "KEYWORD"
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenNumber     TokenType = "NUMBER"
	TokenFloat      TokenType = "FLOAT"
	TokenString     TokenType = "STRING"
	TokenOperator   TokenType = "OPERATOR"
	TokenBoolOp     TokenType = "BOOL_OPERATOR"
	TokenSeparator  TokenType = "SEPARATOR"
	TokenEOF        TokenType = "EOF"
	TokenUnknown    TokenType = "UNKNOWN"
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

type Lexer struct {
	input       string
	position    int
	line        int
	column      int
	currentChar rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 1}
	l.advance()
	return l
}

func (l *Lexer) advance() {
	if l.position < len(l.input) {
		l.currentChar = rune(l.input[l.position])
		l.position++
		l.column++
		if l.currentChar == '\n' {
			l.line++
			l.column = 1
		}
	} else {
		l.currentChar = 0 // EOF
	}
}

func (l *Lexer) peekChar() rune {
	if l.position < len(l.input) {
		return rune(l.input[l.position])
	}
	return 0
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
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
	keywords := map[string]TokenType{
		"if":     TokenKeyword,
		"else":   TokenKeyword,
		"for":    TokenKeyword,
		"func":   TokenKeyword,
		"int":    TokenKeyword, // Added 'int' keyword
		"float":  TokenKeyword, // Added 'float' keyword
		"string": TokenKeyword, // Added 'string' keyword
	}

	if tokenType, isKeyword := keywords[value]; isKeyword {
		return Token{Type: tokenType, Value: value, Line: l.line, Column: l.column}
	}
	return Token{Type: TokenIdentifier, Value: value, Line: l.line, Column: l.column}
}

func (l *Lexer) handleNumber() Token {
	start := l.position - 1
	isFloat := false

	for isDigit(l.currentChar) || (l.currentChar == '.' && !isFloat) {
		if l.currentChar == '.' {
			isFloat = true
		}
		l.advance()
	}

	value := l.input[start : l.position-1]
	if isFloat {
		return Token{Type: TokenFloat, Value: value, Line: l.line, Column: l.column}
	}
	return Token{Type: TokenNumber, Value: value, Line: l.line, Column: l.column}
}

func (l *Lexer) handleString() Token {
	l.advance() // skip opening "
	start := l.position - 1

	for l.currentChar != '"' && l.currentChar != 0 {
		l.advance()
	}

	value := l.input[start : l.position-1]
	l.advance() // skip closing "

	return Token{Type: TokenString, Value: value, Line: l.line, Column: l.column}
}

func (l *Lexer) skipLineComment() {
	for l.currentChar != '\n' && l.currentChar != 0 {
		l.advance()
	}
}

func (l *Lexer) skipBlockComment() {
	for l.currentChar != 0 {
		if l.currentChar == '*' && l.peekChar() == '/' {
			l.advance()
			l.advance()
			break
		}
		l.advance()
	}
}

func (l *Lexer) GetNextToken() Token {
	for l.currentChar != 0 {
		if unicode.IsSpace(l.currentChar) {
			l.advance()
			continue
		}

		if l.currentChar == '/' {
			if l.peekChar() == '/' {
				l.skipLineComment()
				continue
			} else if l.peekChar() == '*' {
				l.skipBlockComment()
				continue
			}
		}

		if isLetter(l.currentChar) {
			return l.handleIdentifier()
		}

		if isDigit(l.currentChar) {
			return l.handleNumber()
		}

		if l.currentChar == '"' {
			return l.handleString()
		}

		switch l.currentChar {
		case '+', '-', '*', '/', '=', '<', '>', ';', ',', '(', ')', '{', '}':
			op := string(l.currentChar)
			l.advance()
			if op == ";" || op == "," || op == "(" || op == ")" || op == "{" || op == "}" {
				return Token{Type: TokenSeparator, Value: op, Line: l.line, Column: l.column}
			}
			return Token{Type: TokenOperator, Value: op, Line: l.line, Column: l.column}
		}

		// Unrecognized character
		fmt.Printf("Unrecognized character at Line %d, Column %d: %q\n", l.line, l.column, l.currentChar)
		l.advance()
		return Token{Type: TokenUnknown, Value: string(l.currentChar), Line: l.line, Column: l.column}
	}

	return Token{Type: TokenEOF, Value: "", Line: l.line, Column: l.column}
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	lexer := NewLexer(input)

	outputFile, err := os.Create("./tokens.txt")
	if err != nil {
		log.Fatalf("Failed to create tokens file: %s", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for {
		token := lexer.GetNextToken()
		fmt.Fprintf(writer, "Type: %s, Value: %s, Line: %d, Column: %d\n", token.Type, token.Value, token.Line, token.Column)
		if token.Type == TokenEOF {
			break
		}
	}
	writer.Flush()

	fmt.Println("Tokens:")
	tokensFile, err := os.Open("./tokens.txt")
	if err != nil {
		log.Fatalf("Failed to open tokens.txt: %s", err)
	}
	defer tokensFile.Close()

	tokensScanner := bufio.NewScanner(tokensFile)
	for tokensScanner.Scan() {
		fmt.Println(tokensScanner.Text())
	}

	if err := tokensScanner.Err(); err != nil {
		log.Fatalf("Error reading tokens file: %s", err)
	}
}
