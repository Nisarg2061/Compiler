import re

class ThreeAddressCodeGenerator:
    def __init__(self):
        self.temp_count = 0
        self.quadruples = []
        self.triples = []
        self.indirect_triples = []
        self.temp_map = {}
        self.temp_indices = []

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

        # Generate a new temporary variable name for the result
        result = self.new_temp()

        # Generate quadruple with T1, T2 style temporaries
        self.quadruples.append((op, left, right, result))

        # Generate triples using the temporary variable directly for assignment
        self.triples.append((op, left, right))

        # Store the index of this triple in the indirect triples list
        index = len(self.triples) - 1
        self.indirect_triples.append(index)

        # Append the temporary variable result to the values stack
        values.append(result)

        # Track temp variable mappings for debugging
        self.temp_map[result] = f"{left} {op} {right}"

        # Store the index of the temporary variable (index for indirect triples)
        self.temp_indices.append(result)

    def to_3ac(self, expression):
        """Convert an expression to 3AC."""
        var, expr = expression.split('=')
        var = var.strip()
        expr = expr.strip()

        # Skip the declaration keywords like 'int', 'float', 'string'
        if expr.startswith('int') or expr.startswith('float') or expr.startswith('string'):
            return  # Ignore lines with these keywords

        tokens = re.findall(r'[a-zA-Z0-9]+|[()+\-*/]', expr)
        values = []
        operators = []

        for token in tokens:
            if token.isalnum():
                values.append(token)
            elif token == '(':
                operators.append(token)
            elif token == ')':
                while operators and operators[-1] != '(':
                    self.apply_operator(operators, values)
                operators.pop()
            else:
                while (operators and operators[-1] != '(' and
                       self.precedence(operators[-1]) >= self.precedence(token)):
                    self.apply_operator(operators, values)
                operators.append(token)

        while operators:
            self.apply_operator(operators, values)

        # Final result for the expression, assign it to the variable
        result = values.pop()

        # Quadruple assignment with the actual temp variable
        self.quadruples.append(('=', result, '', var))

        # Triples assignment directly using the temp variable (e.g., T1)
        self.triples.append(('=', result, ''))

        # Indirect triples assignment using the index (last entry), but without parentheses
        index = len(self.triples) - 1
        self.indirect_triples.append(index)

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
        for i, index in enumerate(self.indirect_triples):
            op, arg1, arg2 = self.triples[index]
            # In indirect triples, use the index values from temp_indices
            indirect_arg1 = self.temp_indices.index(arg1) if arg1 in self.temp_indices else arg1
            indirect_arg2 = self.temp_indices.index(arg2) if arg2 in self.temp_indices else arg2
            print(f"{i}\t{op}\t\t({indirect_arg1})\t({indirect_arg2})")

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
                if expression:
                    expressions.append(expression)

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
