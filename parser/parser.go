package parser

import (
	"cyclik/lexer"
	"go/token"
)

type Parser struct {
	lex *lexer.Lexer //просто ссылка на лексер

	curToken  token.Token // текущий символ
	peekToken token.Token // следующий
}
