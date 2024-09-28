package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("code.txt")
	check(err)
	defer file.Close()

	var code []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code = append(code, scanner.Text())
	}

	joinedCode := strings.Join(code, " ")
	tokens := separateTokens(joinedCode)

	lexicalAnalysis(tokens)
}

func separateTokens(code string) []string {
	delimiters := []rune{'=', '*', '+', '-', '/', '>', '<', '!', ';'}
	tokens := []string{}
	token := ""

	for _, char := range code {
		if unicode.IsSpace(char) {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			continue
		}

		isDelimiter := false
		for _, d := range delimiters {
			if char == d {
				if token != "" {
					tokens = append(tokens, token)
					token = ""
				}
				tokens = append(tokens, string(char))
				isDelimiter = true
				break
			}
		}
		if !isDelimiter {
			token += string(char)
		}
	}
	if token != "" {
		tokens = append(tokens, token)
	}

	return tokens
}

func lexicalAnalysis(tokens []string) {
	keywords := []string{}
	identifiers := []string{}
	constants := []string{}
	delimiters := []string{}
	operators := []string{}

	for _, token := range tokens {
		if token == "int" || token == "float" || token == "var" {
			keywords = append(keywords, token)
		} else if isConstant(token) {
			constants = append(constants, token)
		} else if token == ";" {
			delimiters = append(delimiters, token)
		} else if token == "+" || token == "-" || token == "=" || token == "!=" || token == ">" || token == "<" || token == "*" || token == "/" {
			operators = append(operators, token)
		} else {
			identifiers = append(identifiers, token)
		}
	}

	fmt.Printf("Keywords: %v\n", keywords)
	fmt.Printf("Identifiers: %v\n", identifiers)
	fmt.Printf("Constants: %v\n", constants)
	fmt.Printf("Delimiters: %v\n", delimiters)
	fmt.Printf("Operators: %v\n", operators)
}

func isConstant(token string) bool {
	for _, char := range token {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

