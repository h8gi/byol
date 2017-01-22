package main

import (
	"./lispy"
	"fmt"
	"github.com/chzyer/readline"
)

func display(p *lispy.Parser, line string) {
	p.SetString(line)
	tokens, _ := p.ReadTokens()
	// display innput
	fmt.Println("Lexer------------")
	for i, token := range tokens {
		fmt.Printf("%d: %+v\n", i, token)
		if token.Text == "quit" {
			return
		}
	}
	fmt.Println("------------Lexer")
}

func repl() {
	fmt.Println("Lispy Version 0.0.0.0.1")
	fmt.Println("Press Ctrl+c to Exit")
	rl, err := readline.New("lispy> ")
	if err != nil {
		panic(err)
	}
	p := lispy.Parser{}
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		display(&p, line)
		program, err := p.ParseString(line)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		for _, ast := range program {
			fmt.Println(ast)
		}
		obj, env := lispy.EvalProgram(program, lispy.Environment{})
		fmt.Printf("%v, %v\n", obj, env)
	}
}

func filetest() {
	p := lispy.Parser{}
	ast, err := p.ParseFile("test.scm")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ast)
}

func main() {
	// Print version
	repl()
}
