package main

import (
	"fmt"
)

func main() {
	i := make([]int, 3)
	i[0] = 1
	i[1] = 2
	fmt.Print(len(i), cap(i))
}