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
	fmt.Println(list.Eq(&list))
	list2 := list
	fmt.Println(list.Eq(&list2))
	fmt.Println(lispy.LispNull.Eq(&lispy.LispNull))
	fmt.Printf("%v eq %v = %v\n", syma, symb, syma.Eq(&symb))
	fmt.Printf("%v eq %v = %v\n", syma, syma, syma.Eq(&syma))
	syma2 := lispy.NewSymbol("a")
	fmt.Printf("%v eq %v = %v\n", syma, syma2, syma.Eq(&syma2))

}
