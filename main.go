package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"goscript/ast"
	"goscript/parser"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		byteContent, _ := os.ReadFile(fileName)
		stringCon := string(byteContent)
		p := parser.NewParser(stringCon)
		env := ast.NewEvaluator()
		nodes := p.Parse()
		for _, node := range nodes {
			node.Evaluate(env)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		env := ast.NewEvaluator()
		for {
			fmt.Print("> ")
			scanner.Scan()
			line := scanner.Text()
			if line == "exit" {
				break
			}
			P := parser.NewParser(line)
			nodes := P.Parse()
			for _, node := range nodes {
				node.Evaluate(env)
			}
		}
	}
}
