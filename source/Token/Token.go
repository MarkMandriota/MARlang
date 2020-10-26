package Token

type (
	Token struct {
		Type int
		Val  []byte
		Pos  int
	}
)

const (
	EOF     = iota // 0
	ILLEGAL        // 23..324
	IDENT          // a123, name ...
	FLOAT          // 1.432, 554.32553 ...
	INT            // 2832, 593, 20, 3 ...
	STRING         // "Hello, World!"

	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %

	EQU  // =
	LAR  // >
	LES  // <
	NOT  // !
	LARe // >=
	LESe // <=
	NOTe // !=
	AND  // &
	OR   // |

	ASS // :

	LPAREN   // (
	RPAREN   // )
	LBRACE   // {
	RBRACE   // }
	LBRACKET // [
	RBRACKET // ]

	SEMICOLON // ;
	COMMA     // ,
)
