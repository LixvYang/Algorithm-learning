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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j, k := 0, 0, 0
	pre, cur := 0, 0
	n1, n2 := len(nums1), len(nums2)
	m := n1 + n2
	mid := m / 2
	for k <= mid {
			pre = cur
			if i < n1 && j < n2 {
					if nums1[i] < nums2[j] {
							cur = nums1[i]
							i++
					} else {
							cur = nums2[j]
							j++
					}
			} else if i < n1 {
					cur = nums1[i]
					i++
			} else {
					cur = nums2[j]
					j++
			}
			k++
	}
	if m % 2 == 0 {
			return float64(pre + cur) / 2
	}
	return float64(cur)

}