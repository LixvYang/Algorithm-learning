// 乘积最大子数组
package leetcode


func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 0 {
		return 0
	}

	maxValue, minValue := nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		maxValue_ := max(maxValue*nums[i], max(minValue*nums[i], nums[i]))
		minValue = min(maxValue*nums[i], min(minValue*nums[i], nums[i]))
		maxValue = maxValue_
		ans = max(ans, maxValue)
	}

	return ans
}

