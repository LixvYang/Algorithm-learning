package leetcode

func LargestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1)
	heights = append(heights, 0)	
	ln := len(heights)
	res := 0

	for i := 0; i < ln; i++ {
		for len(stack)>1 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			left := stack[len(stack)-1]
			res = max(res, (i-left-1)*heights[top])
		}
		stack = append(stack, i)
	}
	return res
}


func max(a, b int) int {
	if a > b {
			return a
	}
	return b
}