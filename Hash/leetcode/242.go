package leetcode 

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
			return false
	}

	arr := [26]int{}

	for _, v := range s {
			arr[v - 'a']++
	}
	for _, v := range t {
			arr[v - 'a']--
	}

	for _, v := range arr {
			if v != 0 {
					return false
			}
	}
	return true
}