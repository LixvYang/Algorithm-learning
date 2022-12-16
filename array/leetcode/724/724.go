package leetcode    

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
			sum += v
	}
	leftSum := 0
	rightSum := 0
	for i := 0; i < len(nums); i++ {
			leftSum += nums[i]
			rightSum = sum - leftSum + nums[i]
			if leftSum == rightSum {
					return i
			}
	}
	return -1
}