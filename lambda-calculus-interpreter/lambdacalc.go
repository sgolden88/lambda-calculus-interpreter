package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sgolden88/lambda-calculus-interpreter/ast"
	lex "github.com/sgolden88/lambda-calculus-interpreter/lexer"
	parse "github.com/sgolden88/lambda-calculus-interpreter/parser"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(">>")
	env := NewEnvironment()
	for reader.Scan() { //used this because windows has a two line return and linux has a one line return(\r\n vs \n)and this truncates it
		text := reader.Text()
		if text == "quit()" {
			break
		}
		lexer := lex.NewLexer(text)
		parser := parse.NewParser(&lexer)
		program, err := parser.Parse()
		if err != nil {
			fmt.Println(err)
			print(">>")
			continue
		}

		final := eval(program, env)

		if final != nil {
			println("Final value of expression", final.ExprString())
		} else {
			println()
		}

		print(">>")
	}
}
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
func NewEnvironment() *Environment {
	s := make(map[ast.Identifier]ast.Expression)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[ast.Identifier]ast.Expression
	outer *Environment
}

func (e *Environment) Get(name ast.Identifier) (ast.Expression, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Environment) Set(name ast.Identifier, val ast.Expression) ast.Expression {
	e.store[name] = val
	return val
}
func eval(root ast.AstNode, env *Environment) ast.Expression { //eval issue where same name bindings get replaced when they shouldn't in an application
	switch node := root.(type) {
	case ast.Abstraction:
		temp, ok := env.Get(node.Arg)
		env.Set(node.Arg, node.Arg) //free variables with same name of the argument do not get replaced in the expression
		ret := ast.Abstraction{Arg: node.Arg, Exp: eval(node.Exp, env)}
		if ok {
			env.Set(node.Arg, temp)
		}
		return ret
	case ast.Application:
		left := eval(node.Fun, env)
		if left, isabs := left.(ast.Abstraction); isabs {
			//evaluate abstraction with bound argument
			//because maps are shallow copies always for no good reason besides fuck you
			newenv := NewEnclosedEnvironment(env)
			newenv.Set(left.Arg, eval(node.Arg, env))
			return eval(left.Exp, newenv)

		}
		return ast.Application{left, eval(node.Arg, env)}
	case ast.Identifier:
		if val, ok := env.Get(node); ok { //if identifier is bound, return its bound value, otherwise return the identifier
			return val
		}
		return node
	case ast.Binding:
		val := eval(node.Val, env)
		env.Set(node.Id, val)
		return val
	}
	return nil
}
