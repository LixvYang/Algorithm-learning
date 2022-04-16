package leetcode

func reverseWords(s string) string {
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
			fastIndex++
	}

	for ; fastIndex < len(b); fastIndex++ {
			if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
					continue
			}
			b[slowIndex] = b[fastIndex]
			slowIndex++
	}

	if slowIndex - 1 >0 && b[slowIndex-1] == ' ' {
			b = b[:slowIndex-1]
	} else {
			b = b[:slowIndex]
	}

	reverse(&b, 0, len(b)-1)
	i := 0;
	for i < len(b) {
			j := i
			for ; j < len(b) && b[j] != ' '; j++ {
	}
	reverse(&b, i, j-1)
	i = j
	i++
	}
	return string(b)
}

func reverse(b *[]byte, left, right int) {
for left < right {
	(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
	left++
	right--
}
}

func reverseWords(s string) string {
	b := []byte(s)
	var reverse func(start,end int)

	reverse = func(start, end int){
		for start < end {
			b[start], b[end] = b[end], b[start]
			start += 1
			end -= 1
		}
	}
	reverse(0,len(b)-1)
	start := 0
	end := 0
	i := 0
	res := ""
	for i < len(b) {
		if b[i] != ' '{
			start = i
			for i < len(b) && b[i] != ' '{
				i += 1
			}
			end = i - 1
			reverse(start, end)
			res = res + " " + string(b[start:end+1])
		}
		i += 1
	}
	return res[1:]
}
