package main

import (
	"fmt"
)


func main() {
	main := "abcd227facac"
	pattern := "ac"
	fmt.Println(bfsearch(main, pattern))
}

func bfsearch(main, pattern string) int {
	
	//defensive
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return -1
	}

	for i := 0; i <= len(main) - len(pattern); i++ {
		subStr := main[i:i+len(pattern)]
		if subStr == pattern{
			return i
		}
	}
	return -1
}