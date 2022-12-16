package leetcode

import "testing"

func Test_SmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SmallerNumbersThanCurrent(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
