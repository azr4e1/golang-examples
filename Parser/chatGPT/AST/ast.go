package main

type ASTNode interface{}

type NumberNode struct {
	Value int
}

type BinaryOpNode struct {
	Left     ASTNode
	Operator TokenType
	Right    ASTNode
}
