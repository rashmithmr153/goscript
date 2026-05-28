package ast

type Evaluator struct {
	Store     map[string]any
	Functions map[string]*FunctionDef
	Parent    *Evaluator
}
type FunctionDef struct {
	Params []string
	Body   []Node
}

func (E *Evaluator) Get(name string) (any, bool) {
	if val, exist := E.Store[name]; exist {
		return val, true
	}
	if E.Parent != nil {
		return E.Parent.Get(name)
	}
	return 0, false
}

func (E *Evaluator) Set(name string, value any) {
	E.Store[name] = value
}

func (E *Evaluator) LookupFunction(name string) (*FunctionDef, bool) {
	if fn, exist := E.Functions[name]; exist {
		return fn, true
	}
	if E.Parent != nil {
		return E.Parent.LookupFunction(name)
	}
	return nil, false
}

func NewEvaluator() *Evaluator {
	return &Evaluator{
		Store:     map[string]any{},
		Functions: map[string]*FunctionDef{},
	}
}
func NewEnclosedEvaluator(parent *Evaluator) *Evaluator {
	env := NewEvaluator()
	env.Parent = parent
	return env
}
