package leetcode

import "testing"

func TestNextGreaterElement(t *testing.T) {
	tests := []struct{
		nums1 []int
		nums2 []int
		want []int
	}{
		{[]int{4,1,2}, []int{1,3,4,2}, []int{-1,3,-1}},
		{[]int{2,4}, []int{1,2,3,4}, []int{3,-1}},
	}

	for _, test := range tests {
		if got := NextGreaterElement(test.nums1, test.nums2); !eql(got, test.want) {
			t.Errorf("nextGreaterElement(%v, %v) = %v, want %v", test.nums1, test.nums2, got, test.want)
		}
	}
}

func eql(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}