package parser

import (
	"fmt"
	"goscript/ast"
	"goscript/lexer"
	"strconv"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

func NewParser(input string) Parser {
	lexer := lexer.NewLexer(input)
	return Parser{
		lexer:     lexer,
		curToken:  lexer.NextToken(),
		peekToken: lexer.NextToken(),
	}
}

func (P *Parser) advance() {
	P.curToken = P.peekToken
	P.peekToken = P.lexer.NextToken()
}

func (P *Parser) ParseLetStatement() *ast.LetStatement {
	P.advance() // move past "let"

	name := P.curToken.Value // now curToken is IDENTIFIER
	P.advance()              // move past identifier

	P.advance() // move past "="
	value := P.ParseComparison()
	if P.curToken.Type == lexer.SEMICOLON {
		P.advance()
	}

	return &ast.LetStatement{Name: name, Value: value}
}

func (P *Parser) ParseIfStatement() *ast.IfStatement {
	P.advance()
	condition := P.ParseComparison()
	cons := P.ParseBlock()
	alt := []ast.Node{}
	if P.curToken.Type == lexer.KEYWORD && P.curToken.Value == "else" {
		alt = P.ParseBlock()
	}

	return &ast.IfStatement{
		Condition:   condition,
		Consequence: cons,
		Alt:         alt,
	}
}

func (P *Parser) ParseForStatement() *ast.Forstatement {
	P.advance() //skip for
	condn := P.ParseComparison()
	consq := P.ParseBlock()
	return &ast.Forstatement{
		Condition:   condn,
		Consequence: consq,
	}
}

func (P *Parser) ParseReturnStatement() *ast.ReturnStatement {
	P.advance() //advance "return"
	val := P.ParseComparison()
	if P.curToken.Type == lexer.SEMICOLON {
		P.advance()
	}
	return &ast.ReturnStatement{
		Value: val,
	}
}

func (P *Parser) ParseFactor() ast.Node {
	if P.curToken.Type == lexer.INT {
		val, _ := strconv.Atoi(P.curToken.Value)
		numNode := ast.NumberNode{
			Value: val,
		}
		P.advance()
		return &numNode
	}
	if P.curToken.Type == lexer.STRING {
		strNode := &ast.StringNode{Value: P.curToken.Value}
		P.advance()
		return strNode
	}
	if P.curToken.Type == lexer.KEYWORD &&
		(P.curToken.Value == "true" || P.curToken.Value == "false") {
		val := P.curToken.Value == "true"
		P.advance()
		return &ast.BoolNode{Value: val}
	}
	if P.curToken.Type == lexer.IDENTIFIER {
		name := P.curToken.Value
		P.advance()
		if P.curToken.Type == lexer.LPAREN {
			return P.ParseFuncCall(name)
		}
		return &ast.IdentifierNode{Name: name}
	}
	if P.curToken.Type == lexer.LPAREN {
		P.advance()
		result := P.ParseExpression()
		P.advance()
		return result
	}
	panic(fmt.Sprintf("unexpected token '%s'", P.curToken.Value))
}
func (P *Parser) ParseFuncDef() *ast.FunctionDefNode {
	P.advance() // skip "func"

	name := P.curToken.Value // function name
	P.advance()              // skip name

	P.advance() // skip "("

	params := []string{}
	for P.curToken.Type != lexer.RPAREN {
		if P.curToken.Type == lexer.COMMA {
			P.advance() // skip comma
		} else {
			prm := P.curToken.Value
			params = append(params, prm)
			P.advance()
		}
	}
	P.advance() // skip ")"

	body := P.ParseBlock()
	return &ast.FunctionDefNode{Name: name, Params: params, Body: body}
}

func (P *Parser) ParseFuncCall(name string) *ast.FunctionCallNode {
	P.advance() // skip "("
	args := []ast.Node{}
	for P.curToken.Type != lexer.RPAREN {
		if P.curToken.Type != lexer.COMMA {
			arg := P.ParseComparison()
			args = append(args, arg)
		} else {
			P.advance() // skip comma
		}
	}
	P.advance() // skip ")"
	return &ast.FunctionCallNode{Name: name, Args: args}
}

func (P *Parser) ParseComparison() ast.Node {
	left := P.ParseExpression()
	for P.curToken.Type == lexer.OPERATOR &&
		(P.curToken.Value == "==" || P.curToken.Value == "!=" || P.curToken.Value == ">" || P.curToken.Value == "<" || P.curToken.Value == ">=" || P.curToken.Value == "<=") {
		op := P.curToken.Value
		P.advance()
		right := P.ParseExpression()
		left = &ast.BinaryNode{
			Left:  left,
			Op:    op,
			Right: right,
		}
	}
	return left
}

func (P *Parser) ParseTerm() ast.Node {
	left := P.ParseFactor()
	for P.curToken.Type == lexer.OPERATOR &&
		(P.curToken.Value == "*" || P.curToken.Value == "/") {
		op := P.curToken.Value
		P.advance()
		right := P.ParseFactor()
		left = &ast.BinaryNode{
			Left:  left,
			Op:    op,
			Right: right,
		}
	}
	return left
}
func (P *Parser) ParseExpression() ast.Node {
	left := P.ParseTerm()
	for P.curToken.Type == lexer.OPERATOR &&
		(P.curToken.Value == "+" || P.curToken.Value == "-") {
		op := P.curToken.Value
		P.advance()
		right := P.ParseTerm()
		left = &ast.BinaryNode{
			Left:  left,
			Op:    op,
			Right: right,
		}
	}
	return left
}
func (P *Parser) parseStatement() ast.Node {
	if P.curToken.Type == lexer.KEYWORD && P.curToken.Value == "let" {
		return P.ParseLetStatement()
	}
	if P.curToken.Type == lexer.KEYWORD && P.curToken.Value == "if" {
		return P.ParseIfStatement()
	}
	if P.curToken.Type == lexer.KEYWORD && P.curToken.Value == "for" {
		return P.ParseForStatement()
	}
	if P.curToken.Type == lexer.KEYWORD && P.curToken.Value == "return" {
		return P.ParseReturnStatement()
	}
	if P.curToken.Type == lexer.KEYWORD && P.curToken.Value == "func" {
		return P.ParseFuncDef()
	}
	if P.curToken.Type == lexer.IDENTIFIER && P.peekToken.Type == lexer.LPAREN {
		name := P.curToken.Value
		P.advance()
		node := P.ParseFuncCall(name)
		if P.curToken.Type == lexer.SEMICOLON {
			P.advance()
		}
		return node
	}
	if P.curToken.Type == lexer.SEMICOLON {
		P.advance()
		return nil
	}
	P.advance()
	return nil
}
func (P *Parser) ParseBlock() []ast.Node {
	result := []ast.Node{}
	P.advance()
	for P.curToken.Type != lexer.RBRACE && P.curToken.Type != lexer.EOF {
		node := P.parseStatement()
		if node != nil {
			result = append(result, node)
		}
	}
	P.advance()
	return result
}
func (P *Parser) Parse() []ast.Node {
	result := []ast.Node{}
	for P.curToken.Type != lexer.EOF {
		node := P.parseStatement()
		if node != nil {
			result = append(result, node)
		}
	}
	return result
}
