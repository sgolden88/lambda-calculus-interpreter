package ast

type AstNode interface {
	node()
}
type Expression interface {
	//Cannot use type constraints so that i dont boof shit, have to use function and potentially field constraints
	ExprString() string
	node()
}
type Binding struct {
	Id  Identifier
	Val Expression
}

func (b Binding) node() {

}

type Identifier string

func (i Identifier) ExprString() string {
	return string(i)
}
func (i Identifier) node() {
}

type Abstraction struct {
	Arg Identifier
	Exp Expression
}

func (a Abstraction) ExprString() string {
	return "(λ" + string(a.Arg) + "." + a.Exp.ExprString() + ")"
}
func (a Abstraction) node() {
}

type Application struct {
	Fun Expression
	Arg Expression
}

func (a Application) node() {
}
func (a Application) ExprString() string {
	return "(" + a.Fun.ExprString() + " " + a.Arg.ExprString() + ")"
}
