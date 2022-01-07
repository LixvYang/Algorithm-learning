package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for _, v := range nums1 {
			m[v] = 1
	}

	var res []int

	for _, v := range nums2 {
			if count, ok := m[v]; ok && count > 0 {
					res = append(res, v)
					m[v]--
			}
	}
	return res
}