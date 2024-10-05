package main

import (
	"fmt"
)

func main() {
	input := "if x + 5 else y * 10"
	lexer := NewLexer(input)

	fmt.Println("Tokens:")
	for {
		token := lexer.GetNextToken()
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
		if token.Type == TokenEOF {
			break
		}
	}
}
