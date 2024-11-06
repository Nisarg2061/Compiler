import os
import sys
from enum import Enum

class TokenType(Enum):
    KEYWORD = "KEYWORD"
    IDENTIFIER = "IDENTIFIER"
    NUMBER = "NUMBER"
    OPERATOR = "OPERATOR"
    EOF = "EOF"

class Token:
    def __init__(self, token_type, value):
        self.type = token_type
        self.value = value

    def __repr__(self):
        return f"Token(type={self.type}, value={self.value})"

class SemanticAnalyzer:
    def __init__(self, tokens):
        self.tokens = tokens
        self.symbol_table = {}

    def analyze(self):
        for token in self.tokens:
            if token.type == TokenType.KEYWORD:
                # Handle keyword-related analysis here
                pass
            elif token.type == TokenType.IDENTIFIER:
                self.handle_identifier(token)
            elif token.type == TokenType.OPERATOR:
                self.handle_operator(token)
            elif token.type == TokenType.NUMBER:
                self.handle_number(token)

    def handle_identifier(self, token):
        """Check if the identifier is already declared."""
        if token.value in self.symbol_table:
            print(f"Warning: Variable '{token.value}' is already declared.")
        else:
            self.symbol_table[token.value] = 'Declared'

    def handle_operator(self, token):
        """Handle operator-related semantic rules."""
        print(f"Operator encountered: {token.value}")

    def handle_number(self, token):
        """Handle number-related semantic rules."""
        try:
            value = float(token.value)
            print(f"Number encountered: {value}")
        except ValueError:
            print(f"Error: '{token.value}' is not a valid number.")

def read_tokens_from_file(file_path):
    tokens = []

    # Check if the file exists
    if not os.path.exists(file_path):
        print(f"Error: {file_path} does not exist.")
        sys.exit(1)

    with open(file_path, 'r') as file:
        lines = file.readlines()
        for line in lines:
            # Parse each line to extract token type and value
            parts = line.strip().split(", ")
            if len(parts) == 2:
                try:
                    token_type = TokenType[parts[0].split(": ")[1].strip().upper()]
                    value = parts[1].split(": ")[1].strip()
                    tokens.append(Token(token_type, value))
                except KeyError:
                    print(f"Error: Invalid token type '{parts[0]}' found in line: {line.strip()}")
                except IndexError:
                    print(f"Error: Malformed line: {line.strip()}")
            else:
                print(f"Error: Malformed line: {line.strip()}")

    return tokens

def main():
    # Path to the tokens.txt file
    tokens_file_path = './tokens.txt'

    # Read tokens from the file
    tokens = read_tokens_from_file(tokens_file_path)

    # Perform semantic analysis using the tokens
    analyzer = SemanticAnalyzer(tokens)
    analyzer.analyze()

    # Output semantic analysis result
    print("Semantic analysis completed successfully!")

if __name__ == "__main__":
    main()
