package other

import "fmt"

func convert(m, n int) string {
	// sign check m is a negative
	var sign bool
	var res string
	if m < 0 {
		sign = true
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
		fmt.Println(string(b))
		res = string(b) + res
		fmt.Println(res)
		m /= n
	}
	if sign {
		res = "-" + res
	}
	return res
}

func main() {
	convert(10, 2)
	// fmt.Println(convert(10, 2))
}
