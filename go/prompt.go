package main

import(
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
	
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		
		// display innput
		fmt.Println(line)
	}
}
