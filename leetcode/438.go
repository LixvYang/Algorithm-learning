// 找到字符串中所有字母异位词e
package leetcode

func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	if m > n {
		return nil
	}

	cnt := [26]int{}
	for i := range p {
		cnt[p[i]-'a']--
	}

	var lo, hi int
	ret := make([]int, 0)
	for hi < len(s) {
		x := s[hi] - 'a'
		cnt[x]++

		for cnt[x] > 0 {
			cnt[s[lo]-'a']--
			lo++
		}
		if hi-lo+1 == m {
			ret = append(ret, lo)
			cnt[s[lo]-'a']--
			lo++
		}
		hi++
	}

	return ret
}