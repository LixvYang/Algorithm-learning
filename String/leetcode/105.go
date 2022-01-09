package leetcode

func reverseWords(s string) string {
	b := []byte(s)
	var reverse3 func(start, end int)
	reverse3 = func(start, end int)  {
		for start <= end {
			b[start], b[end] = b[end], b[start]
			start++
			end--
		}
	}

	reverse3(0, len(b)-1)
	start := 0
	end := 0
	i := 0
	res := ""
	for ;i < len(b);i++ {
		if b[i] != ' '{
			start = i
			for i < len(b) && b[i] != ' ' {
				i++
			}
			end = i-1
			reverse3(start, end)
			res = res + " " + string(b[start:end+1])
		}
	}
	return res[1:]
}