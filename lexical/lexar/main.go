package main

func isDelimiter(ch rune) bool  {
  if ch == ' ' || ch == '+' || ch == '-' || ch == '*' || 
        ch == '/' || ch == ',' || ch == ';' || ch == '>' || 
        ch == '<' || ch == '=' || ch == '(' || ch == ')' || 
        ch == '[' || ch == ']' || ch == '{' || ch == '}' {
          return true
        } else {
          return false
        }
}

func isOperator(ch rune) bool  {
  if ch == '+' || ch == '-' || ch == '*' || 
        ch == '/' || ch == '>' || ch == '<' || 
        ch == '=' {
          return true
        }else {
          return false
        }
}

func main()  {
  code := "int i = 0;"
  
}
