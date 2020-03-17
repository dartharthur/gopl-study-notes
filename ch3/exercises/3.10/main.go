package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	lead := n % 3
	if lead == 0 {
		lead = 3
	}

	buf.WriteString(s[:lead])
	for i := lead; i < n-1; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

/* naive first attempt

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i := len(s); i >= 0; i -= 3 {
		if i-3 <= 0 {
			sofar := buf.String()
			buf.Reset()
			buf.WriteString(s[:i] + sofar)
			break
		}
		sofar := buf.String()
		buf.Reset()
		buf.WriteString("," + s[i-3:i] + sofar)
	}

	return buf.String()
}

*/
