// 下一个更大元素1
package leetcode

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	// 初始化res = -1
	for i := range nums1 {
		res[i] = -1
	}
	// mp可通过val找到key
	mp := make(map[int]int)
	for k, v := range nums1 {
		mp[v] = k
	}

	// 单调栈， 装nums2从小到大排序
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			if index, ok := mp[nums2[top]]; ok {
				res[index] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}