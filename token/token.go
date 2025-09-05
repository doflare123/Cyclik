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

	//всякие знаки
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

	// ключевые слова
	FUNCTION    = "FUNCTION"
	RETURN      = "RETURN"
	INT_TYPE    = "INT"
	FLOAT_TYPE  = "FLOAT"
	STRING_TYPE = "STRING"
	BOOL_TYPE   = "BOOL"
	ENTER       = "ENTER"
	PRINT       = "PRINT"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	COMMENT = "COMMENT"
)
