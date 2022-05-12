package leetcode

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