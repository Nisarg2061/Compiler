package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the sample.txt file
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Read the file contents
	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + " " // Concatenate each line with a space
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Create the lexer using the input from the file
	lexer := NewLexer(input)

	// Print the tokens
	fmt.Println("Tokens:")
	for {
		token := lexer.GetNextToken()
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
		if token.Type == TokenEOF {
			break
		}
	}
}

