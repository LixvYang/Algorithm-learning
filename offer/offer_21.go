// 调整数组顺序使奇数位于偶数前面
package offer

func exchange(nums []int) []int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
			if nums[lo]%2 == 0 && nums[hi]%2!=0 {
					nums[lo], nums[hi] = nums[hi], nums[lo]
			}

			for lo< len(nums)&& nums[lo]%2!=0 {
					lo++
			}

			for hi>=0 && nums[hi]%2 == 0 {
					hi--
			}
	}
	return nums
}