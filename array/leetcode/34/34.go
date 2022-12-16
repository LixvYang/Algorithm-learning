package leetcode

func searchRange(nums []int, target int) []int {
	leftBorder := getLeftBorder(nums, target)
	rightBorder := getRightBorder(nums, target)

	// 情况一：target 在数组范围的右边或者左边，例如数组{3, 4, 5}，target为2或者数组{3, 4, 5},target为6，此时应该返回{-1, -1}
	// 情况二：target 在数组范围中，且数组中不存在target，例如数组{3,6,7},target为5，此时应该返回{-1, -1}
	// 情况三：target 在数组范围中，且数组中存在target，例如数组{3,6,7},target为6，此时应该返回{1, 1}
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	// 情况三
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	// 情况二
	return []int{-1, -1}
}

func getLeftBorder(nums []int, target int) int {
	leftBorder := -2
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] >= target {
			right = mid - 1
			leftBorder = right
		} else {
			left = mid + 1
		}
		// fmt.Println(left, right)
	}
	return leftBorder
}

func getRightBorder(nums []int, target int) int {
	rightBorder := -2

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] <= target {
			left = mid + 1
			rightBorder = left
		} else {
			right = mid - 1
		}
	}
	return rightBorder
}
