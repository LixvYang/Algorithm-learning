package leetcode

func reverseStr(s string, k int) string {
	ss := []byte(s)
	len := len(s)

	for i := 0; i < len; i += 2 * k {
			if i + k <= len {
					reverse(ss[i:i+k])
			} else {
					reverse(ss[i:len])
			}
	}
	return string(ss)
}

func reverse(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
			s[left], s[right] = s[right], s[left]
			left++
			right--
	}
}