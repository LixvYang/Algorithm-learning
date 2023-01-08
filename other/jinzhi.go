package other

import "fmt"

func convert(m, n int) string {
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
		res = string(b) + res
		m /= n
	}
	if sign {
		res = "-" + res
	}
	return res
}

func main() {
	fmt.Println(convert(10, 2))
}
