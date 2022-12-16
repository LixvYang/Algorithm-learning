package leetcode

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	l, r := 0, len(arr)-1

	for l < len(arr)-1 && arr[l] < arr[l+1] {
		l++
	}

	for r > 0 && arr[r] < arr[r-1] {
		r--
	}

	if l == r && l != 0 && r != len(arr)-1 {
		return true
	}
	return false
}
