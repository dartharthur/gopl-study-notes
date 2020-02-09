// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// LineData represents metadata for a given line of text.
type LineData struct {
	count int
	// files FileSet
}

// FileSet represents the set of files in which a line of text appears.
// type FileSet struct {
// 	filename string
// }

func main() {
	lineData := make(map[string]*LineData)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lineData)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, lineData)
			f.Close()
		}
	}
	for line, lineData := range lineData {
		if lineData.count > 1 {
			fmt.Printf("%d\t%s\n", lineData.count, line)
		}
	}
}

func countLines(f *os.File, lineData map[string]*LineData) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		_, ok := lineData[input.Text()]
		if !ok {
			lineData[input.Text()] = &LineData{1}
		}
		lineData[input.Text()].count++
	}
	// NOTE: ignoring potential errors from input.Err()
}
