package main

import (
	"./lispy"
	"fmt"
	"github.com/chzyer/readline"
)

func display(lx *lispy.Lexer, line string) {
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

func main() {
	// Print version
	fmt.Println("Lispy Version 0.0.0.0.1")
	fmt.Println("Press Ctrl+c to Exit")
	rl, err := readline.New("lispy> ")
	if err != nil {
		panic(err)
	}
	lx := lispy.Lexer{}

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		display(&lx, line)
		lx.SetString(line)
		lx.ReadToken()
		ast, err := lx.Datum()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(ast)
	}
}
