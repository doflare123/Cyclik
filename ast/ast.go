package ast

import "go/token"

type Node interface { //база
	TokenLiteral() string
}

type Statement interface { //инструкции
	Node
	statementNode()
}

type Expression interface { //вычисления
	Node
	expressionNode()
}

type PrintStatment struct {
	Token token.Token
	Value Expression
}
