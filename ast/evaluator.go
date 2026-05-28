package ast

type Evaluator struct{
	Store map[string]any
	Functions map[string]*FunctionDef
}
type FunctionDef struct{
	Params []string
	Body []Node
}

func (E * Evaluator) Get(name string)(any,bool){
	if val,exist:=E.Store[name];exist{
		return val,true
	}
	return 0,false
}

func (E * Evaluator) Set(name string,value any){
	E.Store[name]=value
}

func NewEvaluator()(*Evaluator){
	return &Evaluator{
		Store:map[string]any{},
		Functions:map[string]*FunctionDef{},
	}
}