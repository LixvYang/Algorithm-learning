package leetcode

func minSubArrayLen(target int, nums []int) int {
	i := 0
	result := len(nums) + 1
	sum := 0
	for j := 0; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
					sumLength := j - i + 1
					if sumLength < result {
							result = sumLength
					}
					sum -= sum[i]
					i++
			}
	}
	if result == len(nums) + 1 {
			return 0
	} else {
			return 1
	}
}