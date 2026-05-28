package ast

import "fmt"

type Node interface {
	Evaluate(env *Evaluator) (any, bool)
}

func toInt(v any) int {
	if i, ok := v.(int); ok {
		return i
	}
	return 0
}

func toBool(v any) bool {
	switch val := v.(type) {
	case bool:
		return val
	case int:
		return val != 0
	case string:
		return val != ""
	default:
		return false
	}
}

type NumberNode struct {
	Value int
}

type BinaryNode struct {
	Left  Node
	Op    string
	Right Node
}
type ReturnStatement struct {
	Value Node
}

func (r *ReturnStatement) Evaluate(env *Evaluator) (any, bool) {
	val, _ := r.Value.Evaluate(env)
	return val, true
}

type FunctionDefNode struct {
	Name   string
	Params []string
	Body   []Node
}

func (f *FunctionDefNode) Evaluate(env *Evaluator) (any, bool) {
	env.Functions[f.Name] = &FunctionDef{
		Params: f.Params,
		Body:   f.Body,
	}
	return 0, false
}

type FunctionCallNode struct {
	Name string
	Args []Node
}

func (F *FunctionCallNode) Evaluate(env *Evaluator) (any, bool) {
	fn, exists := env.Functions[F.Name]
	if !exists {
		if F.Name == "print" {
			if len(F.Args) == 0 {
				panic("print() requires at least one argument")
			}
			val, _ := F.Args[0].Evaluate(env)
			fmt.Println(val)
			return 0, false
		}
		panic(fmt.Sprintf("undefined function '%s'", F.Name))
	}
	temEnv := NewEvaluator()
	for i, param := range fn.Params {
		val, _ := F.Args[i].Evaluate(env)
		temEnv.Set(param, val)
	}
	var result any = 0
	var isReturn bool
	for _, node := range fn.Body {
		result, isReturn = node.Evaluate(temEnv)
		if isReturn {
			return result, false
		}
	}
	return result, false
}

type IfStatement struct {
	Condition   Node
	Consequence []Node
	Alt         []Node
}

func (If *IfStatement) Evaluate(env *Evaluator) (any, bool) {
	condnRes, isReturn := If.Condition.Evaluate(env)
	var result any
	if toBool(condnRes) {
		for _, node := range If.Consequence {
			result, isReturn = node.Evaluate(env)
			if isReturn {
				return result, true
			}
		}
		return result, isReturn
	}
	for _, node := range If.Alt {
		result, isReturn = node.Evaluate(env)
		if isReturn {
			return result, true
		}
	}
	return result, false
}

type Forstatement struct {
	Condition   Node
	Consequence []Node
}

func (Fr *Forstatement) Evaluate(env *Evaluator) (any, bool) {
	condRes, isReturn := Fr.Condition.Evaluate(env)
	var result any = 0
	for toBool(condRes) {
		for _, node := range Fr.Consequence {
			result, isReturn = node.Evaluate(env)
			if isReturn {
				return result, true
			}
		}
		condRes, isReturn = Fr.Condition.Evaluate(env)
	}
	return result, false
}

type IdentifierNode struct {
	Name string
}

func (Id *IdentifierNode) Evaluate(env *Evaluator) (any, bool) {
	if val, exist := env.Get(Id.Name); exist {
		return val, false
	}
	return 0, false
}

type LetStatement struct {
	Name  string
	Value Node
}

func (ls *LetStatement) Evaluate(env *Evaluator) (any, bool) {
	val, _ := ls.Value.Evaluate(env)
	env.Set(ls.Name, val)
	return val, false
}

type StringNode struct {
	Value string
}

func (s *StringNode) Evaluate(env *Evaluator) (any, bool) {
	return s.Value, false
}

type BoolNode struct {
	Value bool
}

func (b *BoolNode) Evaluate(env *Evaluator) (any, bool) {
	return b.Value, false
}

func (Nd *NumberNode) Evaluate(env *Evaluator) (any, bool) {
	return Nd.Value, false
}

func (Bd *BinaryNode) Evaluate(env *Evaluator) (any, bool) {
	lv, _ := Bd.Left.Evaluate(env)
	rv, _ := Bd.Right.Evaluate(env)
	leftval := toInt(lv)
	RightVal := toInt(rv)
	opertor := Bd.Op
	result := 0
	switch opertor {
	case "+":
		result = leftval + RightVal
	case "*":
		result = leftval * RightVal
	case "==":
		if leftval == RightVal {
			result = 1
		} else {
			result = 0
		}
	case "!=":
		if leftval != RightVal {
			result = 1
		} else {
			result = 0
		}
	case ">":
		if leftval > RightVal {
			result = 1
		} else {
			result = 0
		}
	case "<":
		if leftval < RightVal {
			result = 1
		} else {
			result = 0
		}
	case ">=":
		if leftval >= RightVal {
			result = 1
		} else {
			result = 0
		}
	case "<=":
		if leftval <= RightVal {
			result = 1
		} else {
			result = 0
		}
	default:
		result = 0
	}
	return result, false
}
