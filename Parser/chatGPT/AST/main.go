package main

import (
	"encoding/json"
	_ "fmt"
	"os"
)

func evaluate(node ASTNode) int {
	switch n := node.(type) {
	case *NumberNode:
		return n.Value
	case *BinaryOpNode:
		leftVal := evaluate(n.Left)
		rightVal := evaluate(n.Right)
		switch n.Operator {
		case TokenPlus:
			return leftVal + rightVal
		case TokenMinus:
			return leftVal - rightVal
		case TokenMul:
			return leftVal * rightVal
		case TokenDiv:
			return leftVal / rightVal
		}
	}
	panic("invalid AST node")
}

func main() {
	lexer := NewLexer("3 + 5 * (10 - 4) / 24")
	parser := NewParser(lexer)
	ast := parser.expr()

	f, err := os.Create("ast.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(ast)
	if err != nil {
		panic(err)
	}
	// result := evaluate(ast)
	// fmt.Println(result) // Output should be 18
}
