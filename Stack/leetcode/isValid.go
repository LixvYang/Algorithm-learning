func isValid(s string) bool {
	n := len(s)

	if n % 2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
        ']': '[',
        '}': '{',
	}

	stack := []byte{}

	for i:=0;i<n;i++ {
		if  pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack) - 1] != pairs[s[i]] {
				return false
			}
			stack  = stack[:len(stack)-1]
		} else {
			stack = append(stack,s[i]}
		}
	}
	return len(stack) == 0
}

func isValid(s string) bool {
	var stack []byte
	var top int = -1
	for _,v := range s{
		if v == '('{
			stack = append(stack,')')
			top++
		} else if v == '{' {
			stack = append(stack,'}')
			top++
		} else if v == '[' {
			stack = append(stack,']')
			top++
		}else{
			if top > -1 && byte(v) == stack[top]{
				stack = stack[:top]
				top--
			} else {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}