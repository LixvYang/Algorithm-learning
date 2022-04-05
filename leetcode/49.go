// 字母异位词
package leetcode

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {return s[i] < s[j]})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	res := make([][]string, 0)
	for _, v := range mp {
		res = append(res, v)
	}

	return res
}