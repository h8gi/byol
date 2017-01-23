package main

import (
	"./lispy"
	"fmt"
)

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
	lispy.Repl()
}
