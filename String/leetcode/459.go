// 重复的子字符串
package leetcode

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	next := make([]int, n)
	getNext(next, s)
	
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 0;i < len(s); i++ {
		for j>0 && s[i] != s[j] {
			j = next[j]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}