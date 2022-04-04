// 接雨水
package leetcode

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	maxLeft, maxRight := make([]int, len(height)), make([]int, len(height))
	maxLeft[0], maxRight[len(height)-1] = height[0], height[len(height)-1]

	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	for i := len(height)-2; i>=0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	} 

	sum := 0
	for i := 0; i < len(height); i++ {
		count := min(maxLeft[i], maxRight[i]) - height[i]
		if count > 0 {
			sum += count
		}
	}
	return sum
}
