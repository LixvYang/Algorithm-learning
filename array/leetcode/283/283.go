package leetcode

func moveZeroes(nums []int) {
	slowIndex, fastIndex := 0, 0

	for ; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != 0 {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}

	for i := slowIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
