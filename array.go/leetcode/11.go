package leetcode

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	ans := 0
	for l < r {
		w := r-l
		h := 0
		if height[l] > height[r] {
			h = height[r]
			r--
		} else {
			h = height[l]
			l++
		}

		tmp := w * h
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}