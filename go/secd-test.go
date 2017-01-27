package main

import (
	"./lispy"
	"fmt"
)

func main() {
	syma, _ := lispy.NewSymbol("a")
	symb, _ := lispy.NewSymbol("b")
	fmt.Println(lispy.NewList(syma, symb, symb))
}
