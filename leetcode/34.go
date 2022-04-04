// 在排序数组中查找元素的第一个和最后一个位置
package leetcode

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid+1
		} else {
			r = mid-1
		}
	}

	if l == 0 || nums[l-1] != target {
		return []int{-1, -1}
	}

	cnt := 0
  for i := l+1; i < len(nums); i++ {
      if nums[i] == target {
          cnt++
      }
  }
  return []int{l, l+cnt}
}