package parser

import (
	"goscript/ast"
	"goscript/lexer"
	"strconv"
)

type Parser struct{
	lexer *lexer.Lexer
	curToken lexer.Token
	peekToken lexer.Token
}

func NewParser(input string) (Parser){
	lexer:=lexer.NewLexer(input)
	return Parser{
		lexer: lexer,
		curToken: lexer.NextToken(),
		peekToken: lexer.NextToken(),
	}
}

func (P *Parser)advance(){
	P.curToken=P.peekToken
	P.peekToken=P.lexer.NextToken()
}

func (P *Parser) ParseLetStatement() *ast.LetStatement {
    P.advance() // move past "let"

    name := P.curToken.Value // now curToken is IDENTIFIER
    P.advance() // move past identifier

    P.advance() // move past "="
    value := P.ParseExpression()
    P.advance() // move past number

    return &ast.LetStatement{Name: name, Value: value}
}

func (P *Parser)ParseFactor() ast.Node{
	if P.curToken.Type==lexer.INT{
		val,_:=strconv.Atoi(P.curToken.Value)
		numNode:=ast.NumberNode{
			Value:val ,
		}
		P.advance()
		return &numNode
	}
	if P.curToken.Type == lexer.IDENTIFIER {
		IdNode:=ast.IdentifierNode{
			Name: P.curToken.Value,
		}
		P.advance()
		return &IdNode
	}
	if P.curToken.Type==lexer.LPAREN{
		P.advance()
		result:=P.ParseExpression()
		P.advance()
		return result
	}
	return nil
}

func (P *Parser)ParseTerm() ast.Node{
	left:=P.ParseFactor()
	for P.curToken.Type==lexer.OPERATOR &&
	(P.curToken.Value=="*"||P.curToken.Value=="/"){
		op:=P.curToken.Value
		P.advance()
		right:=P.ParseFactor()
		left=&ast.BinaryNode{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}
func (P *Parser)ParseExpression() ast.Node{
	left:=P.ParseTerm()
	for P.curToken.Type==lexer.OPERATOR &&
	(P.curToken.Value=="+"||P.curToken.Value=="-"){
		op:=P.curToken.Value
		P.advance()
		right:=P.ParseTerm()
		left=&ast.BinaryNode{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}