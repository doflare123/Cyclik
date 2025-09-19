package ast

// абстрактное синтаксическое дерево (AST), которое превращает код в набор структурированных объектов

import "cyclik/token"

// гл интер программы
type Program struct {
	Statements []Statement // хранение всех инструкций верхнего уровня
}

// возвращает литерал первого токена программы
func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) > 0 {
		return prog.Statements[0].TokenLiteral()
	}
	return ""
}

// база проги
type Node interface {
	TokenLiteral() string
}

// интерфейс для инструкций
type Statement interface {
	Node
	statementNode()
}

// Содержимое функций и фиг скобок
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (block *BlockStatement) statementNode() {}
func (block *BlockStatement) TokenLiteral() string {
	return block.Token.Literal
}

//  интерфейс для выражений (арифметика, идентификаторы, вызовы)
type Expression interface {
	Node
	expressionNode()
}

// инструкция, которая состоит из одного выражения
// Например вызов функции без присваивания результата
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// вызов функции
type CallExpression struct {
	Token     token.Token
	Function  Expression   //кого вызываем
	Arguments []Expression //список аргументов
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

// отдельная инструкция печати
type PrintStatment struct {
	Token token.Token
	Value Expression
}

// объявление функции
// Содержит её имя, параметры, тело и возвращаемый тип
type FunctionStatment struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
	ReturnType token.TokenType
}

func (funcStat *FunctionStatment) statementNode() {}
func (funcStat *FunctionStatment) TokenLiteral() string {
	return funcStat.Token.Literal
}

// объявление переменной
type VarStatment struct {
	Token token.Token     // токен типа
	Name  *Identifier     // имя переменной
	Type  token.TokenType // тип переменной
	Value Expression      // инициализации (по желанию)
}

func (varStat *VarStatment) statementNode() {}
func (varStat *VarStatment) TokenLiteral() string {
	return varStat.Token.Literal
}

// узел представляющий имя переменной или функции
type Identifier struct {
	Token token.Token
	Value string
}

func (ind *Identifier) expressionNode() {}
func (ind *Identifier) TokenLiteral() string {
	return ind.Token.Literal
}

// узел для целых чисел
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (intLit *IntegerLiteral) expressionNode() {}
func (intLit *IntegerLiteral) TokenLiteral() string {
	return intLit.Token.Literal
}

// инструкция возврата значения из функции
type ReturnStatment struct {
	Token       token.Token
	ValueReturn Expression
}

func (retStat *ReturnStatment) statementNode() {}
func (retStat *ReturnStatment) TokenLiteral() string {
	return retStat.Token.Literal
}

// унарное выражение перед операндом
// вроде ! перед переменными/значениями
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// бинарное выражение с оператором
// a + b, x * y.
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
