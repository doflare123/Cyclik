package lexer

import "cyclik/token"

var keywords = map[string]token.TokenType{
	"int":     token.INT_TYPE,
	"float32": token.FLOAT_TYPE,
	"string":  token.STRING_TYPE,
	"bool":    token.BOOL_TYPE,
	"return":  token.RETURN,
	"enter":   token.ENTER,
	"true":    token.TRUE,
	"false":   token.FALSE,
	"CD":      token.PRINT,
	"F":       token.FUNCTION,
}

func LookIndent(ident string) token.TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return token.IDENT
}
