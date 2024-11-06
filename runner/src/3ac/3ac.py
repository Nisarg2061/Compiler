import re

class ThreeAddressCodeGenerator:
    def __init__(self):
        self.temp_count = 0
        self.quadruples = []
        self.triples = []
        self.indirect_triples = []
        self.indirect_table = []
        self.temp_map = {}  # Dictionary to store temp variable mappings

    def new_temp(self):
        """Generate a new temporary variable name."""
        self.temp_count += 1
        return f"T{self.temp_count}"

    def precedence(self, op):
        """Return precedence of operators."""
        if op in ('+', '-'):
            return 1
        if op in ('*', '/'):
            return 2
        return 0

    def apply_operator(self, operators, values):
        """Apply an operator to two values and generate TAC."""
        op = operators.pop()
        right = values.pop()
        left = values.pop()
        result = self.new_temp()

        # Generate quadruple
        self.quadruples.append((op, left, right, result))
        self.triples.append((op, left, right))
        index = len(self.triples) - 1
        self.indirect_table.append(index)
        self.indirect_triples.append((op, left, right))

        # Track the temp variable mappings
        self.temp_map[result] = f"{left} {op} {right}"

        values.append(result)

    def to_3ac(self, expression):
        """Convert an expression to 3AC."""
        # Extract the variable and expression part: x = a * b + c * (a * B)
        var, expr = expression.split('=')
        var = var.strip()
        expr = expr.strip()

        tokens = re.findall(r'[a-zA-Z0-9]+|[()+\-*/]', expr)
        values = []
        operators = []

        for token in tokens:
            if token.isalnum():  # If it's an operand (variable or number)
                values.append(token)
            elif token == '(':
                operators.append(token)
            elif token == ')':
                while operators and operators[-1] != '(':
                    self.apply_operator(operators, values)
                operators.pop()  # Pop '('
            else:  # Operator
                # Handle precedence correctly by applying operators
                while (operators and operators[-1] != '(' and
                       self.precedence(operators[-1]) >= self.precedence(token)):
                    self.apply_operator(operators, values)
                operators.append(token)

        # Apply the remaining operators
        while operators:
            self.apply_operator(operators, values)

        # Assign the result to the variable
        result = values.pop()
        self.quadruples.append(('=', result, '', var))
        self.triples.append(('=', result, ''))
        index = len(self.triples) - 1
        self.indirect_table.append(index)
        self.indirect_triples.append(('=', result, ''))

        # Track the assignment in the temp variable map
        self.temp_map[var] = result

    def print_tables(self):
        """Print the generated TAC tables."""
        print("\nQuadruples Table:")
        print("Operator\tArg1\tArg2\tResult")
        for (op, arg1, arg2, res) in self.quadruples:
            print(f"{op}\t\t{arg1}\t{arg2}\t{res}")

        print("\nTriples Table:")
        print("Index\tOperator\tArg1\tArg2")
        for i, (op, arg1, arg2) in enumerate(self.triples):
            print(f"{i}\t{op}\t\t{arg1}\t{arg2}")

        print("\nIndirect Triples Table:")
        print("Index\tOperator\tArg1\tArg2")
        for i in self.indirect_table:
            op, arg1, arg2 = self.indirect_triples[i]
            print(f"{i}\t{op}\t\t{arg1}\t{arg2}")

        print("\nTemporary Variables Mapping:")
        for temp_var, expression in self.temp_map.items():
            print(f"{temp_var} = {expression}")

# Main loop to read expressions from sample.txt
def main():
    generator = ThreeAddressCodeGenerator()

    # Read expressions from sample.txt
    try:
        with open('sample.txt', 'r') as file:
            lines = file.readlines()
            expressions = []
            for line in lines:
                expression = line.strip()
                if expression:  # Avoid empty lines
                    expressions.append(expression)

            # Process each expression
            for expression in expressions:
                try:
                    generator.to_3ac(expression)
                except ValueError as e:
                    print(f"Error processing expression '{expression}': {e}")

    except FileNotFoundError:
        print("sample.txt file not found.")

    # Print all tables
    generator.print_tables()

# Run the program
if __name__ == "__main__":
    main()
