// Echo2 prints its command-line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	fmt.Println("invoked by: ", os.Args[0])
	for i, arg := range os.Args[1:] {
		fmt.Printf("at index %v with arg %v\n", i, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
