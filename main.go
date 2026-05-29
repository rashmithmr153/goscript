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
func safeParse(input string) (nodes []ast.Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("parse error: %v", r)
		}
	}()
	p := parser.NewParser(input)
	nodes = p.Parse()
	return nodes, nil
}

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		byteContent, _ := os.ReadFile(fileName)
		nodes, err := safeParse(string(byteContent))
		if err != nil {
			fmt.Println(err)
			return
		}
		env := ast.NewEvaluator()
		if err := safeRun(nodes, env); err != nil {
			fmt.Println(err)
		}
	} else {
		// REPL mode
		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); err != nil {
			fmt.Println("error in scanner: ", err)
			return
		}
		env := ast.NewEvaluator()
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			if line == "" {
				continue
			}
			if line == "exit" {
				break
			}
			nodes, err := safeParse(line)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err := safeRun(nodes, env); err != nil {
				fmt.Println(err)
			}
		}
	}
}
