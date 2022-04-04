// 寻找两个正序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// Language: go
// Path: leetcode\5.go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res []int
	m, n := len(nums1), len(nums2)
	var l1, l2 int

	for l1 < m && l2 < n {
			if nums1[l1] < nums2[l2] {
					res = append(res, nums1[l1])
					l1++
			} else {
					res = append(res, nums2[l2])
					l2++
			}
	}

	res = append(res, nums1[l1:]...)
	res = append(res, nums2[l2:]...)

	length := m+n
	if length%2 == 1 {
			return float64(res[length/2])
	}

	m1 := res[length/2]
	m2 := res[length/2-1]
	return float64(m1+m2)/2.0
}