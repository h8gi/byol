package main

import (
	"./lispy"
	"fmt"
	"github.com/chzyer/readline"
)

func display(lx *lispy.Parser, line string) {
	lx.SetString(line)
	tokens, _ := lx.ReadTokens()
	// display innput
	for i, token := range tokens {
		fmt.Print(i, token, "\n")
		if token.Text == "quit" {
			return
		}
	}
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
		ast, err := p.ParseString(line)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(ast)
	}
}

func main() {
	// Print version
	p := lispy.Parser{}
	ast, err := p.ParseFile("test.scm")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ast)
}
