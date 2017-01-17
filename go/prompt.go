package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	// Print version
	fmt.Println("Lispy Version 0.0.0.0.1\n")
	fmt.Println("Press Ctrl+c to Exit\n")
	reader := bufio.NewReaderSize(os.Stdin, 2048)
	for {
		// output our prompt
		fmt.Print("lispy> ")

		// read input
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}

		// display innput
		fmt.Println(string(line))
	}
}
