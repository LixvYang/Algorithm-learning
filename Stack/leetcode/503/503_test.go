package leetcode

import "testing"

func TestNextGreaterElements(t *testing.T) {
	test := []struct{
		nums []int
		want []int
	}{
		{[]int{1,2,1}, []int{2, -1, 2},
		{[]int{1,2,3,4,3}, []int{2,3,4,-1,4}},
	}

	for _, test := range tests {
		if got := NextGreaterElements(test.nums); !eql(nums, test.got) {
			t.Errorf("nextGreaterElement(%v) = %v, want %v", test.nums, got, test.want)
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
