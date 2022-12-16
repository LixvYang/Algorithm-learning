package leetcode941

func SmallerNumbersThanCurrent(nums []int) []int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	quickSort(nums, 0, len(nums)-1)
	m := make(map[int]int)
	for i := len(nums) - 1; i >= 0; i-- {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = m[copyNums[i]]
	}

	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(nums, start, end)
	quickSort(nums, start, i-1)
	quickSort(nums, i+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			if nums[j] != nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
