package ast

import "fmt"

type Node interface {
	Evaluate(env *Evaluator) (int, bool)
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

func (r *ReturnStatement) Evaluate(env *Evaluator) (int, bool) {
	val, _ := r.Value.Evaluate(env)
	return val, true
}

type FunctionDefNode struct {
	Name   string
	Params []string
	Body   []Node
}

func (f *FunctionDefNode) Evaluate(env *Evaluator) (int, bool) {
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

func (F *FunctionCallNode) Evaluate(env *Evaluator) (int, bool) {
	fn, exists := env.Functions[F.Name]
	if !exists {
		if F.Name == "print" {
			val, _ := F.Args[0].Evaluate(env)
			fmt.Println(val)
			return 0, false
		}
		return 0, false
	}
	temEnv := NewEvaluator()
	for i, param := range fn.Params {
		val, _ := F.Args[i].Evaluate(env) // evaluate arg in OUTER env
		temEnv.Set(param, val)            // store in LOCAL env
	}
	result := 0
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

func (If *IfStatement) Evaluate(env *Evaluator) (int, bool) {
	condnRes, isReturn := If.Condition.Evaluate(env)
	result := 0
	if condnRes != 0 {
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

func (Fr *Forstatement) Evaluate(env *Evaluator) (int, bool) {
	condRes, isReturn := Fr.Condition.Evaluate(env)
	result := 0
	for condRes != 0 {
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

func (Id *IdentifierNode) Evaluate(env *Evaluator) (int, bool) {
	if val, exist := env.Get(Id.Name); exist {
		return val, false
	}
	return 0, false
}

type LetStatement struct {
	Name  string
	Value Node
}

func (ls *LetStatement) Evaluate(env *Evaluator) (int, bool) {
	val, _ := ls.Value.Evaluate(env)
	env.Set(ls.Name, val)
	return val, false
}

func (Nd *NumberNode) Evaluate(env *Evaluator) (int, bool) {
	return Nd.Value, false
}

func (Bd *BinaryNode) Evaluate(env *Evaluator) (int, bool) {
	leftval, _ := Bd.Left.Evaluate(env)
	opertor := Bd.Op
	RightVal, _ := Bd.Right.Evaluate(env)
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
