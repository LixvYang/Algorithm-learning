// 承最多水的容器
package leetcode


func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		maxArea = max(Area, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}