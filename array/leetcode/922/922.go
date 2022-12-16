package leetcode

func sortArrayByParityII(nums []int) []int {
	res := make([]int, len(nums))
	evenIndex, oddIndex := 0, 1

	for i := 0; i < len(nums); i++ {
			if nums[i] % 2 == 0 {
					res[evenIndex] = nums[i]
					evenIndex += 2
			} else {
					res[oddIndex] = nums[i]
					oddIndex += 2
			}
	}
	return res
}