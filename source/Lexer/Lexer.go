package Lexer

import (
	. "../Token"
)

type Lexer struct {
	input []byte
	rPos  int
}

func New(input []byte) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) Lex() []*Token {
	var tokens []*Token
	for ; l.rPos < len(l.input); l.rPos++ {
		switch l.input[l.rPos] {
		case ' ', '\n', '\t', 'v', '\r':
			continue
		case '+':
			tokens = append(tokens, &Token{Type: ADD})
		case '-':
			tokens = append(tokens, &Token{Type: SUB})
		case '*':
			tokens = append(tokens, &Token{Type: MUL})
		case '/':
			tokens = append(tokens, &Token{Type: DIV})
		case '%':
			tokens = append(tokens, &Token{Type: MOD})
		case '=':
			tokens = append(tokens, &Token{Type: EQU})
		case '>':
			if l.rPos+1 < len(l.input) && l.input[l.rPos+1] == '=' {
				tokens = append(tokens, &Token{Type: LARe})
				l.rPos++
			} else {
				tokens = append(tokens, &Token{Type: LAR})
			}
		case '<':
			if l.rPos+1 < len(l.input) && l.input[l.rPos+1] == '=' {
				tokens = append(tokens, &Token{Type: LESe})
				l.rPos++
			} else {
				tokens = append(tokens, &Token{Type: LES})
			}
		case '!':
			if l.rPos+1 < len(l.input) && l.input[l.rPos+1] == '=' {
				tokens = append(tokens, &Token{Type: NOTe})
				l.rPos++
			} else {
				tokens = append(tokens, &Token{Type: NOT})
			}
		case '&':
			tokens = append(tokens, &Token{Type: AND})
		case '|':
			tokens = append(tokens, &Token{Type: OR})
		case ':':
			tokens = append(tokens, &Token{Type: ASS})
		case '(':
			tokens = append(tokens, &Token{Type: LPAREN})
		case ')':
			tokens = append(tokens, &Token{Type: RPAREN})
		case '{':
			tokens = append(tokens, &Token{Type: LBRACE})
		case '}':
			tokens = append(tokens, &Token{Type: RBRACE})
		case '[':
			tokens = append(tokens, &Token{Type: LBRACKET})
		case ']':
			tokens = append(tokens, &Token{Type: RBRACKET})
		case '"':
			var str []byte
			for l.rPos++; l.rPos < len(l.input) && l.input[l.rPos] != '"'; l.rPos++ {
				str = append(str, l.input[l.rPos])
			}
			tokens = append(tokens, &Token{Type: STRING, Val: str})
		default:
			if IsLetter(l.input[l.rPos]) {
				tokens = append(tokens, l.ReadIdent())
			} else if IsDigit(l.input[l.rPos]) {
				tokens = append(tokens, l.ReadNum())
			} else {
				tokens = append(tokens, &Token{Type: ILLEGAL})
			}
		}
	}

	return tokens
}

func (l *Lexer) ReadNum() *Token {
	t := new(Token)
	retPos := l.rPos
	t.Type = INT
loop:
	for l.rPos < len(l.input) && IsDigit(l.input[l.rPos]) {
		t.Val = append(t.Val, l.input[l.rPos])
		retPos = l.rPos
		l.rPos++
	}
	if l.rPos+1 < len(l.input) && l.input[l.rPos] == '.' {
		l.rPos++
		if IsDigit(l.input[l.rPos]) && t.Type != FLOAT {
			t.Type = FLOAT
			t.Val = append(t.Val, '.')
			goto loop
		} else {
			t.Type = ILLEGAL
		}
	}
	l.rPos = retPos
	return t
}

func (l *Lexer) ReadIdent() *Token {
	t := new(Token)
	retPos := l.rPos
	t.Type = IDENT
	for l.rPos < len(l.input) && (IsLetter(l.input[l.rPos]) || IsDigit(l.input[l.rPos])) {
		t.Val = append(t.Val, l.input[l.rPos])
		retPos = l.rPos
		l.rPos++
	}
	l.rPos = retPos
	return t
}

func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch == '_'
}
