import os
import sys
from enum import Enum
from typing import List, Dict, Union

class TokenType(Enum):
    KEYWORD = "KEYWORD"
    IDENTIFIER = "IDENTIFIER"
    NUMBER = "NUMBER"
    OPERATOR = "OPERATOR"
    EOF = "EOF"
    SEPARATOR = "SEPARATOR"  # Add SEPARATOR for tokens like semicolons, commas, etc.

class Token:
    def __init__(self, token_type: TokenType, value: str) -> None:
        self.type: TokenType = token_type
        self.value: str = value

    def __repr__(self) -> str:
        return f"Token(type={self.type}, value={self.value})"

class SemanticAnalyzer:
    def __init__(self, tokens: List[Token]) -> None:
        self.tokens: List[Token] = tokens
        self.symbol_table: Dict[str, Union[str, float]] = {}  # Store variable names and their types

    def analyze(self) -> None:
        i = 0
        while i < len(self.tokens):
            token = self.tokens[i]
            if token.type == TokenType.KEYWORD and token.value in ['int', 'float']:
                # Declare variable when a keyword is encountered
                if i + 1 < len(self.tokens) and self.tokens[i + 1].type == TokenType.IDENTIFIER:
                    variable_name = self.tokens[i + 1].value
                    self.declare_variable(variable_name, token.value.lower())
                    i += 1  # Skip the next token (the identifier)
            elif token.type == TokenType.IDENTIFIER:
                self.handle_identifier(token)
            elif token.type == TokenType.OPERATOR:
                self.handle_operator(token)
            elif token.type == TokenType.NUMBER:
                self.handle_number(token)
            i += 1

    def handle_identifier(self, token: Token) -> None:
        """Handle variable identifiers and check if they are declared."""
        print(f"Identifier encountered: {token.value}")
        if token.value not in self.symbol_table:
            print(f"Error: Variable '{token.value}' is used but not declared.")
            return

    def handle_operator(self, token: Token) -> None:
        """Handle operator-related semantic rules."""
        print(f"Operator encountered: {token.value}")
        # Here, we need to ensure operands are declared and of correct type for the operation.
        # Operator handling logic goes here.

    def handle_number(self, token: Token) -> None:
        """Handle number-related semantic rules."""
        print(f"Number encountered: {token.value}")
        try:
            # Try converting the value to a float (assuming all numbers are floats for simplicity)
            value = float(token.value)
            print(f"Valid number: {value}")
        except ValueError:
            print(f"Error: '{token.value}' is not a valid number.")
            return

    def declare_variable(self, variable_name: str, value: str) -> None:
        """Declare a variable with a specific type."""
        if value == "float":
            self.symbol_table[variable_name] = "float"
        elif value == "int":
            self.symbol_table[variable_name] = "int"
        print(f"Declaring new variable: {variable_name} of type '{self.symbol_table[variable_name]}'")

    def check_assignment(self, variable_name: str, value: Union[str, float]) -> None:
        """Check if the value being assigned matches the declared type of the variable."""
        if variable_name not in self.symbol_table:
            print(f"Error: Variable '{variable_name}' has not been declared.")
            return

        expected_type = self.symbol_table[variable_name]

        # Checking if the value being assigned is compatible with the variable's type
        if isinstance(value, float) and expected_type == "int":
            print(f"Error: Cannot assign a float to variable '{variable_name}' of type 'int'.")
        elif isinstance(value, str) and expected_type != "string":
            print(f"Error: Cannot assign a string to variable '{variable_name}' of type '{expected_type}'.")

    def check_operator_compatibility(self, operand1: str, operand2: str, operator: str) -> None:
        """Check if the operands are compatible with the operator."""
        # Check if operands are numbers for arithmetic operations
        operand1_type = self.symbol_table.get(operand1)
        operand2_type = self.symbol_table.get(operand2)

        if operand1_type == "float" and operand2_type == "float":
            print(f"Operator {operator} between {operand1} and {operand2} is valid.")
        elif operand1_type == "string" and operand2_type == "string" and operator == "+":
            print(f"Operator {operator} between {operand1} and {operand2} is valid.")
        else:
            print(f"Error: Invalid operation {operator} between {operand1} ({operand1_type}) and {operand2} ({operand2_type}).")

def read_tokens_from_file(file_path: str) -> List[Token]:
    tokens: List[Token] = []

    # Check if the file exists
    if not os.path.exists(file_path):
        print(f"Error: {file_path} does not exist.")
        sys.exit(1)

    with open(file_path, 'r') as file:
        lines = file.readlines()
        for line in lines:
            # Parsing logic to handle 'Type: <type>, Value: <value>, Line: <line>, Column: <column>' format
            parts = line.strip().split(", ")
            if len(parts) >= 2:
                try:
                    # Extract the type and value
                    token_type_part = parts[0].split(": ")[1].strip().upper()
                    token_value_part = parts[1].split(": ")[1].strip()

                    # Determine the correct token type
                    try:
                        token_type = TokenType[token_type_part]
                    except KeyError:
                        print(f"Error: Invalid token type '{token_type_part}' found in line: {line.strip()}")
                        continue

                    # Add the token to the list
                    tokens.append(Token(token_type, token_value_part))

                except IndexError:
                    print(f"Error: Malformed line: {line.strip()}")
            else:
                print(f"Error: Malformed line: {line.strip()}")

    return tokens

def main() -> None:
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
