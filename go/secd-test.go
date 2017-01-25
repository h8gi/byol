package main

import (
	"./lispy"
	"fmt"
)

func main() {
	fmt.Print(lispy.Stack{lispy.LObj{Type: lispy.Boolean, Value: false}})
}
