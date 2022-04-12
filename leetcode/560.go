// 和为K的子数组
package leetcode

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count

}

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0 : 1}
	preSum := 0

	for num := range nums {
		preSum += num
		if hash[preSum - k] > 0 {
			count += hash[preSum - k]
		}
	}
	return count
}