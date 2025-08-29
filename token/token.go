package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT          = "IDENT"
	INT_LITERAL    = "INT_LITERAL"
	FLOAT_LITERAL  = "FLOAT_LITERAL"
	STRING_LITERAL = "STRING_LITERAL"
	BOOL_LITERAL   = "BOOL_LITERAL"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	STAR   = "*"
	SLASH  = "/"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	INT      = "INT"
	FLOAT    = "FLOAT"
	STRING   = "STRING"
	BOOL     = "BOOL"
	ENTER    = "ENTER"
)
