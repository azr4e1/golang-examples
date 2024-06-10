package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
	lexer  *Lexer
	tokens []Token
	pos    int
}

func NewParser(lexer *Lexer) *Parser {
	tokens := []Token{}
	for {
		token := lexer.NextToken()
		if token.Type == TokenEOF {
			break
		}
		tokens = append(tokens, token)
	}
	return &Parser{lexer: lexer, tokens: tokens, pos: 0}
}

func (p *Parser) currentToken() Token {
	if p.pos >= len(p.tokens) {
		return Token{Type: TokenEOF, Value: ""}
	}
	return p.tokens[p.pos]
}

func (p *Parser) eat(tokenType TokenType) {
	if p.currentToken().Type == tokenType {
		p.pos++
	} else {
		panic(fmt.Sprintf("unexpected token: %s", p.currentToken().Value))
	}
}

func (p *Parser) factor() ASTNode {
	token := p.currentToken()

	switch token.Type {
	case TokenNumber:
		p.eat(TokenNumber)
		value, _ := strconv.Atoi(token.Value)
		return &NumberNode{Value: value}
	case TokenLParen:
		p.eat(TokenLParen)
		result := p.expr()
		p.eat(TokenRParen)
		return result
	default:
		panic(fmt.Sprintf("unexpected token: %s", token.Value))
	}
}

func (p *Parser) term() ASTNode {
	result := p.factor()

	for p.currentToken().Type == TokenMul || p.currentToken().Type == TokenDiv {
		token := p.currentToken()
		if token.Type == TokenMul {
			p.eat(TokenMul)
			result = &BinaryOpNode{
				Left:     result,
				Operator: TokenMul,
				Right:    p.factor(),
			}
		} else if token.Type == TokenDiv {
			p.eat(TokenDiv)
			result = &BinaryOpNode{
				Left:     result,
				Operator: TokenDiv,
				Right:    p.factor(),
			}
		}
	}

	return result
}

func (p *Parser) expr() ASTNode {
	result := p.term()

	for p.currentToken().Type == TokenPlus || p.currentToken().Type == TokenMinus {
		token := p.currentToken()
		if token.Type == TokenPlus {
			p.eat(TokenPlus)
			result = &BinaryOpNode{
				Left:     result,
				Operator: TokenPlus,
				Right:    p.term(),
			}
		} else if token.Type == TokenMinus {
			p.eat(TokenMinus)
			result = &BinaryOpNode{
				Left:     result,
				Operator: TokenMinus,
				Right:    p.term(),
			}
		}
	}

	return result
}
