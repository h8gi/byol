package main

import (
	"./lispy"
	"fmt"
)

func main() {
	p := &lispy.Parser{}
	program, _ := p.ParseString("(a b c) 5")
	obj := program[0]
	fmt.Println(obj)
	obj.Pop()
	fmt.Println(obj)
	obj.Push(obj)
	fmt.Println(obj)
	car, _ := obj.Pop()
	fmt.Println(obj, car)
}
