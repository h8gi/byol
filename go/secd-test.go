package main

import (
	"./lispy"
	"fmt"
)

func main() {
	syma := lispy.NewSymbol("a")
	symb := lispy.NewSymbol("b")
	list := lispy.NewList(syma, symb, syma)
	fmt.Println(list)
	fmt.Println(list.Eq(list))
	list2 := list
	fmt.Println(list.Eq(list2))
	fmt.Println(lispy.LispNull.Eq(lispy.LispNull))
	fmt.Printf("%v eq %v = %v\n", syma, symb, syma.Eq(symb))
	fmt.Printf("%v eq %v = %v\n", syma, syma, syma.Eq(syma))
	syma2 := lispy.NewSymbol("a")
	fmt.Printf("%v eq %v = %v\n", syma, syma2, syma.Eq(syma2))
	str1 := lispy.LObj{Type: lispy.DTString, Value: "FOO"}
	str2 := lispy.LObj{Type: lispy.DTString, Value: "FOO"}
	fmt.Printf("%v eq %v = %v\n", str1, str2, str1.Eq(str2))
	list3 := lispy.NewList(syma, symb, syma)
	fmt.Printf("%v eq %v = %v\n", list, list3, list.Eq(list3))
	a := lispy.Cons(syma, symb)
	b := lispy.Cons(syma, symb)
	fmt.Printf("%v eq %v = %v\n", a, b, a.Eq(b))
	fmt.Printf("%v eq %v = %v\n", a, a, a.Eq(a))
	c := a
	fmt.Printf("%v eq %v = %v\n", a, c, a.Eq(c))
	fmt.Println(lispy.NewVector(syma, syma, symb, syma))
	fmt.Println(list.Length())
	fmt.Println(list)
	list.Pop()
	fmt.Println(list)
	fmt.Println(list2)
}
