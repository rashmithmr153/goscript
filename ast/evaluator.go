package ast

type Evaluator struct{
	Store map[string]int
}

func (E * Evaluator) Get(name string)(int,bool){
	if val,exist:=E.Store[name];exist{
		return val,true
	}
	return 0,false
}

func (E * Evaluator) Set(name string,value int){
	E.Store[name]=value
}