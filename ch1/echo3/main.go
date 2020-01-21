// Echo3 prints its command-line arguments

package main

import (
	"fmt"
	"os"
	"strings"
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
	fmt.Println(strings.Join(input, " "))
	// fmt.Printf("%.5vms elapsed\n", time.Since(start).Milliseconds())
}

func main() {
	echo(os.Args[1:])
}
