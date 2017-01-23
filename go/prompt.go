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
	var line string
	var err error
	var contFlag bool
	fmt.Println("Lispy Version 0.0.0.0.1")
	fmt.Println("Press Ctrl+c to Exit")
	rl, err := readline.New("lispy> ")
	if err != nil {
		panic(err)
	}

	p := lispy.Parser{}

	for {
		tmpline, err := rl.Readline()
		if contFlag {
			line = line + "\n" + tmpline
			contFlag = false
			rl.SetPrompt("lispy> ")
		} else {
			line = tmpline
		}
		if err != nil {
			break
		}
		display(&p, line)
		program, err := p.ParseString(line)
		if err != nil {
			switch err.(type) {
			case *lispy.UnclosedError:
				contFlag = true
				rl.SetPrompt("... ")
				continue
			default:
				fmt.Println(err.Error())
				continue
			}
		}

		for _, ast := range program {
			fmt.Println(ast)
		}
		obj, err := lispy.EvalProgram(program)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("result> %v\n", obj)
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
