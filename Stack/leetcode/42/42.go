package leetcode

func Trap(height []int) int {
	sum := 0
	for i := 0; i < len(height); i++ {
		if i== 0 || i == len(height)-1 {
			continue
		}

		rHeight := height[i]
		lHeight := height[i]
		for r := i+1; r < len(height); r++ {
			if rHeight < height[r] {
				rHeight = height[r]
			}
		}

		for l := i-1; l >= 0; l-- {
			if lHeight < height[l] {
				lHeight = height[l]
			}
		}

		h := min(lHeight, rHeight)-height[i]
		if h>0 {
			sum+=h
		}
	}
	return sum
}

func Trap2(height []int) int {
	sum := 0
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	maxLeft[0] = height[0]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight[len(height)-1] = height[len(height)-1]
	for i := len(height)-2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	} 

	for i := 0; i < len(height); i++ {
		h := min(maxLeft[i], maxRight[i]) - height[i]
		if h > 0 {
			sum += h
		}
	}
	return sum
}

func Trap3(height []int) int {
	sum := 0
	stack := []int{}
	for i := 0; i < len(height); i++ {
			for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					if len(stack)>0 {
							left := stack[len(stack)-1]
							curWidth := i-left-1
							curHeight := min(height[i], height[left]) - height[top]
							sum += curWidth*curHeight
					}
			}

			stack = append(stack, i)
	}
	return sum
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
