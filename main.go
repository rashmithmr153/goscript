package main

import (
	"fmt"
	"goscript/ast"
	"goscript/parser"
)

func main() {
	p := parser.NewParser("let x = 10; let y = x + 5;")
	env := &ast.Evaluator{Store: map[string]int{}}
	stmt1 := p.ParseLetStatement()
	stmt1.Evaluate(env)
	stmt2 := p.ParseLetStatement()
	stmt2.Evaluate(env)
	x,_:=env.Get("x")
	y,_:=env.Get("y")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
