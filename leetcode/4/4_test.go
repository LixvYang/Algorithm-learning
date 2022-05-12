package leetcode

import "testing"

func TestfindMedianSortedArrays(t *testing.T {
	var test = []struct{
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, tt := range test {
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays= %v, want %v", got, tt.want)
			}
	}
}