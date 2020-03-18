package main

import "fmt"

func main() {
	fmt.Println(isAnagram("listen", "silent"))    // true
	fmt.Println(isAnagram("aardvark", "ardvark")) // false
}

func isAnagram(s1, s2 string) bool {
	letters := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		letters[s1[i]]++
	}
	for j := 0; j < len(s2); j++ {
		letters[s2[j]]--
		if letters[s2[j]] < 0 {
			return false
		}
	}
	for _, v := range letters {
		if v != 0 {
			return false
		}
	}
	return true
}
