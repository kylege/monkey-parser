package ast

import (
	"monkey-parser/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}
	return p.Statements[0].TokenLiteral()
}

func (p *Program) String() string {
	var b strings.Builder
	for _, s := range p.Statements {
		b.WriteString(s.String())
	}
	return b.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Value
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) String() string {
	var b strings.Builder
	b.WriteString(l.TokenLiteral() + " ")
	b.WriteString(l.Name.String())
	b.WriteString(" = ")
	if l.Value != nil {
		b.WriteString(l.Value.String())
	}
	b.WriteString(";")
	return b.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
	var b strings.Builder
	b.WriteString(r.TokenLiteral() + " ")
	if r.ReturnValue != nil {
		b.WriteString(r.ReturnValue.String())
	}
	b.WriteString(";")
	return b.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}
func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}

type IntergerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntergerLiteral) expressionNode() {}
func (i *IntergerLiteral) TokenLiteral() string {
	return i.Token.Literal
}
func (i *IntergerLiteral) String() string {
	return i.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p *PrefixExpression) String() string {
	var b strings.Builder
	b.WriteString("(")
	b.WriteString(p.Operator)
	b.WriteString(p.Right.String())
	b.WriteString(")")
	return b.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (i *InfixExpression) String() string {
	var b strings.Builder
	b.WriteString("(")
	b.WriteString(i.Left.String())
	b.WriteString(" " + i.Operator + " ")
	b.WriteString(i.Right.String())
	b.WriteString(")")
	return b.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var b strings.Builder
	for _, s := range bs.Statements {
		b.WriteString(s.String())
	}
	return b.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var b strings.Builder
	b.WriteString("if")
	b.WriteString(ie.Condition.String())
	b.WriteString(" ")
	b.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		b.WriteString("else ")
		b.WriteString(ie.Alternative.String())
	}
	return b.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var b strings.Builder
	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	b.WriteString(fl.TokenLiteral())
	b.WriteString("(")
	b.WriteString(strings.Join(params, ", "))
	b.WriteString(")")
	b.WriteString(fl.Body.String())
	return b.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var b strings.Builder
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	b.WriteString(ce.Function.String())
	b.WriteString("(")
	b.WriteString(strings.Join(args, ","))
	b.WriteString(")")

	return b.String()
}
