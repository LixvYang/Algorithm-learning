package leetcode

func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(s); i++ {
			if b[i] == ' ' {
					result = append(result, []byte("%20")...)
			} else {
					result = append(result, b[i])
			}
	} 
	return string(result)
}

func replaceSpace2(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	for _, v := range b {
			if v == ' ' {
					spaceCount++
			}
	}

	resizeCount := spaceCount*2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)
	i, j := length-1, len(b) - 1
	for i >= 0 {
			if b[i] != ' ' {
					b[j] = b[i]
					i--
					j--
			} else {
					b[j] = '0'
					b[j-1] = '2'
					b[j-2] = '%'
					j-=3
					i--
			}
	}
	return string(b)
}