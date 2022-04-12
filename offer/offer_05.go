// 替换空格
package offer

func replaceSpace(s string) string {
	b := []byte(s)
	spaceCount := 0
	for _, v := range b {
			if v == ' ' {
					spaceCount++
			}
	}

	tmp := make([]byte, len(s)+spaceCount*2)
	i, j := len(b)-1, len(tmp)-1
	for i >= 0 {
			if b[i] != ' ' {
					tmp[j] = b[i]
					i--
					j--
			} else {
					tmp[j] = '0'
					tmp[j-1] = '2'
					tmp[j-2] = '%'
					i--
					j-=3
			}
	}
	return string(tmp)
}