// expr    := term (('+' | '-') term)*
// term    := factor (('*' | '/') factor)*
// factor  := INTEGER | '(' expr ')'

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenNumber
	TokenPlus
	TokenMinus
	TokenMul
	TokenDiv
	TokenLParen
	TokenRParen
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, pos: 0}
}

func (l *Lexer) NextToken() Token {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]

		switch {
		case unicode.IsDigit(rune(ch)):
			start := l.pos
			for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
				l.pos++
			}
			return Token{Type: TokenNumber, Value: l.input[start:l.pos]}

		case ch == '+':
			l.pos++
			return Token{Type: TokenPlus, Value: string(ch)}

		case ch == '-':
			l.pos++
			return Token{Type: TokenMinus, Value: string(ch)}

		case ch == '*':
			l.pos++
			return Token{Type: TokenMul, Value: string(ch)}

		case ch == '/':
			l.pos++
			return Token{Type: TokenDiv, Value: string(ch)}

		case ch == '(':
			l.pos++
			return Token{Type: TokenLParen, Value: string(ch)}

		case ch == ')':
			l.pos++
			return Token{Type: TokenRParen, Value: string(ch)}

		case unicode.IsSpace(rune(ch)):
			l.pos++

		default:
			panic(fmt.Sprintf("unexpected character: %c", ch))
		}
	}

	return Token{Type: TokenEOF, Value: ""}
}

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

func (p *Parser) factor() int {
	token := p.currentToken()

	switch token.Type {
	case TokenNumber:
		p.eat(TokenNumber)
		value, _ := strconv.Atoi(token.Value)
		return value
	case TokenLParen:
		p.eat(TokenLParen)
		result := p.expr()
		p.eat(TokenRParen)
		return result
	default:
		panic(fmt.Sprintf("unexpected token: %s", token.Value))
	}
}

func (p *Parser) term() int {
	result := p.factor()

	for p.currentToken().Type == TokenMul || p.currentToken().Type == TokenDiv {
		token := p.currentToken()
		if token.Type == TokenMul {
			p.eat(TokenMul)
			result *= p.factor()
		} else if token.Type == TokenDiv {
			p.eat(TokenDiv)
			result /= p.factor()
		}
	}

	return result
}

func (p *Parser) expr() int {
	result := p.term()

	for p.currentToken().Type == TokenPlus || p.currentToken().Type == TokenMinus {
		token := p.currentToken()
		if token.Type == TokenPlus {
			p.eat(TokenPlus)
			result += p.term()
		} else if token.Type == TokenMinus {
			p.eat(TokenMinus)
			result -= p.term()
		}
	}

	return result
}

func main() {
	lexer := NewLexer("3 + 5 * (10 - 6) / 2")
	parser := NewParser(lexer)
	result := parser.expr()
	fmt.Println(result) // Output should be 8
}
