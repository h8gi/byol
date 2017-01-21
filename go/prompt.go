package main

import (
	"./lispy"
	"fmt"
	"github.com/chzyer/readline"
)

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
		lx.SetString(line)
		// tokens, err := lx.ReadTokens()
		// // display innput
		// for i, token := range tokens {
		// 	fmt.Print(i, token, "\n")
		// 	if token.Text() == "quit" {
		// 		return
		// 	}
		// }
		p := lispy.Parser{}
		p.SetLexer(&lx)
		if err := p.Lispy(); err != nil {
			fmt.Println(err.Error())
		}
	}
}
