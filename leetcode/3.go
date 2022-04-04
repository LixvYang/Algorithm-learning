// 无重复字符的最长子串
// hashtable
func lengthOfLongestSubstring(s string) int {
	var lo, hi, max int
	m := make(map[byte]int)
	for hi < len(s) {
		if count, ok := m[s[hi]]; ok && count > 0 {
			m[s[lo]]--
			lo++
			continue
		}

		res = max(res, hi-lo+1)
		m[s[hi]]++
		hi++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}