// Echo2 prints its command-line arguments

package main

import (
	"fmt"
	"os"
)

func buildInputData(reps int, input []string) []string {
	var built []string
	for i := 0; i < reps; i++ {
		built = append(built, input...)
	}
	return built
}

func echo(input []string) {
	// start := time.Now()
	s, sep := "", ""
	// fmt.Println("invoked by: ", os.Args[0])
	for _, arg := range input {
		// fmt.Printf("at index %v with arg %v\n", i, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// fmt.Printf("%.5vms elapsed\n", time.Since(start).Milliseconds())
}

func main() {
	echo(os.Args[1:])
}
