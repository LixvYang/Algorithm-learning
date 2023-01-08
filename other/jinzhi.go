package main

import "fmt"

func convert(m, n int) string {
	// check m is a negative
	var negative bool
	var res string
	if m < 0 {
		negative = true
		m = -m
	}
	for m > 0 {
		q := m % n
		var b byte
		if q <= 9 {
			b = '0' + byte(q)
		} else {
			b = 'A' + byte(q-10)
		}
		fmt.Println("b: " + string(b))
		res = string(b) + res
		fmt.Println("res: ", res)
		m /= n
	}
	if negative {
		res = "-" + res
	}
	return res
}

func main() {
	convert(10, 2)
	// fmt.Println(convert(10, 2))
}
