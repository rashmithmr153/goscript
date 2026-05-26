package ast

type Node interface{
	Evaluate(env *Evaluator) int
}

type NumberNode struct{
	Value int
}

type BinaryNode struct{
	Left Node
	Op string
	Right Node
}

type LetStatement struct {
    Name  string
    Value Node
}

type IdentifierNode struct{
	Name string
}

func (Id* IdentifierNode) Evaluate(env *Evaluator)int{
	if val,exist:=env.Get(Id.Name);exist{
		return val
	}
	return 0
}

func (ls *LetStatement) Evaluate(env *Evaluator) int {
    val := ls.Value.Evaluate(env)
    env.Set(ls.Name, val)
    return val
}

func (Nd *NumberNode) Evaluate(env *Evaluator)(int){
	return Nd.Value
}

func (Bd *BinaryNode) Evaluate(env *Evaluator)(int){
	leftval:=Bd.Left.Evaluate(env)
	opertor:=Bd.Op
	RightVal:=Bd.Right.Evaluate(env)
	result:=0
	switch opertor {
	case "+":
		result=leftval+RightVal
	case "*":
		result=leftval*RightVal
	default:
		result=0
	}
	return result
}

