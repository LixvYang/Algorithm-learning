package array

import "math"

func myAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	s = s[i:]
	ans := 0
	flag := false

	for i, v := range s {
		if v >= '0' && v <= '9' {
			ans = ans*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			flag = true
		} else if v == '+' && i == 0 {
		} else {
			break
		}
		if ans > math.MaxInt32 {
			if flag {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	if flag {
		return -1 * ans
	}
	return ans
}
