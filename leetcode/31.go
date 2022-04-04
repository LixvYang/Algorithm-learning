// 下一个排列
package leetcode
// 把左边一个较大的数放到右边，然后从右边开始，找到第一个比左边大的数，然后交换，然后从右边开始，找到第一个比左边大的数，然后交换，直到左边的数都比右边的大
func nextPermutation(nums []int)  {
	i := len(nums) -2
	for i > 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1 
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	l, r := i+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}