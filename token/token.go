package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Q: how comes this doesn't have a name? How do we ensure that when it's updated, the rest of the code is aware of that (for tokens)?
// A: because these constants are part of this package. Refer to them via `token.EOF` for example.
const (
	ILLEGAL		= "ILLEGAL"
	EOF 			= "EOF"
	IDENT			= "IDENT"
	INT				= "INT"
	ASSIGN		= "="
	PLUS			= "+"
	COMMA			= ","
	SEMICOLON	= ";"
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"
	FUNCTION	= "FUNCTION"
	LET				= "LET"
)