package token

type TokenType string

// keywords map of the string literal and the associated TokenType
var keywords = map[string]TokenType{
	"func": FUNCTION,
	"var":  VAR,
}

// LookUpIdentifier returns the TokenType for an associated identifier using the
// keywords map
func LookUpIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENTIFIER
}

// Token represents each (or multiple) characters as one unit
type Token struct {
	Type    TokenType
	Literal string
}

// NewToken creates a new instance of Token
func NewToken(tokenType TokenType, char byte) Token {
	return Token{Type: tokenType, Literal: string(char)}
}

// All available tokens
const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	ASSIGN     = "="
	PLUS       = "+"
	MINUS      = "-"
	BANG       = "!"
	ASTERISK   = "*"
	SLASH      = "/"
	LT         = "<"
	GT         = ">"
	COMMA      = ","
	SEMICOLON  = ";"
	LPAREN     = "("
	RPAREN     = ")"
	LBRACE     = "{"
	RBRACE     = "}"
	FUNCTION   = "FUNCTION"
	VAR        = "VAR"
)
