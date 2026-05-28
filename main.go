package main

import (
	"bufio"
	"fmt"
	"goscript/ast"
	"goscript/parser"
	"os"
)

func safeRun(nodes []ast.Node, env *ast.Evaluator) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("runtime error: %v", r)
		}
	}()
	for _, node := range nodes {
		_, isReturn := node.Evaluate(env)
		if isReturn {
			break
		}
	}
	return nil
}

func main() {
	if len(os.Args) > 1 {
		// file mode
		fileName := os.Args[1]
		byteContent, _ := os.ReadFile(fileName)
		p := parser.NewParser(string(byteContent))
		env := ast.NewEvaluator()
		nodes := p.Parse()
		if err := safeRun(nodes, env); err != nil {
			fmt.Println(err)
		}
	} else {
		// REPL mode
		scanner := bufio.NewScanner(os.Stdin)
		if err:=scanner.Err();err!=nil{
			fmt.Println("error in scanner: ",err)
			return
		}
		env := ast.NewEvaluator()
		for {
			fmt.Print("> ")
			scanner.Scan()
			line := scanner.Text()
			if line == "exit" {
				break
			}
			p := parser.NewParser(line)
			nodes := p.Parse()
			if err := safeRun(nodes, env); err != nil {
				fmt.Println(err)
			}
		}
	}
}
