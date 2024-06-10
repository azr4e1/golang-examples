package main

import (
	"fmt"
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
