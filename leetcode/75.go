// 颜色分类
package leetcode


func sortColors(nums []int)  {
	n := len(nums)
	if n < 2 {
		return 
	}

	p0, p2 := 0, n-1
	for i := 0; i <= p2; {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
			i++
		} else if nums[i] == 1 {
			i++
		}
	}
}