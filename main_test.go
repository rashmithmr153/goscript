package main

import (
	"goscript/ast"
	"goscript/parser"
	"testing"
)

func runProgram(input string) *ast.Evaluator {
	p := parser.NewParser(input)
	env := ast.NewEvaluator()
	nodes := p.Parse()
	for _, node := range nodes {
		_, isReturn := node.Evaluate(env)
		if isReturn {
			break
		}
	}
	return env
}

func TestLetStatement(t *testing.T) {
	env := runProgram("let x = 10;")
	x, _ := env.Get("x")
	if x != 10 {
		t.Errorf("expected x=10, got %d", x)
	}
}

func TestElseBlock(t *testing.T) {
	env := runProgram("let x = 5; if x == 10 { let y = 1; } else { let y = 2; }")
	y, _ := env.Get("y")
	if y != 2 {
		t.Errorf("expected y=2, got %d", y)
	}
}

func TestArithmeticInCondition(t *testing.T) {
	env := runProgram("let x = 2 * 5; if x == 10 { let y = 1; }")
	y, _ := env.Get("y")
	if y != 1 {
		t.Errorf("expected y=1, got %d", y)
	}
}

func TestVariableReference(t *testing.T) {
	env := runProgram("let x = 10; let z = x + 5; if z == 15 { let y = 99; }")
	y, _ := env.Get("y")
	if y != 99 {
		t.Errorf("expected y=99, got %d", y)
	}
}

func TestConditionFalse(t *testing.T) {
	env := runProgram("let x = 10; if x > 20 { let y = 1; } else { let y = 0; }")
	y, _ := env.Get("y")
	if y != 0 {
		t.Errorf("expected y=0, got %d", y)
	}
}

func TestNestedIf(t *testing.T) {
	env := runProgram("let x = 10; if x > 5 { let y = 1; if y == 1 { let z = 99; } }")
	z, _ := env.Get("z")
	if z != 99 {
		t.Errorf("expected z=99, got %d", z)
	}
}
func TestForLoop(t *testing.T) {
	env := runProgram("let x = 0; for x < 5 { let x = x + 1; }")
	x, _ := env.Get("x")
	if x != 5 {
		t.Errorf("expected x=5, got %d", x)
	}
}
func TestReturnStatement(t *testing.T) {
	env := runProgram("let x = 10; return x;")
	x, _ := env.Get("x")
	if x != 10 {
		t.Errorf("expected x=10, got %d", x)
	}
}
func TestFunction(t *testing.T) {
	env := runProgram(`
        func add(x, y) {
            return x + y;
        }
        let z = add(10, 5);
    `)
	z, _ := env.Get("z")
	if z != 15 {
		t.Errorf("expected z=15, got %d", z)
	}
}
