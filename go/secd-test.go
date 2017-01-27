package main

import (
	"./lispy"
	"fmt"
)

func main() {
	syma, _ := lispy.NewSymbol("A")
	symb, _ := lispy.NewSymbol("B")
	env := lispy.NewEnv()
	env.Define(syma, symb)
	env.Define(symb, syma)
	env.Define(symb, symb)
	fmt.Println(env)
	result, err := env.LookUp(syma)
	fmt.Println(result, err)
	result, err = env.LookUp(symb)
	fmt.Println(result, err)
	fmt.Println(lispy.InitialEnv())
}
