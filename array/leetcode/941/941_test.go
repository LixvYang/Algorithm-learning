package leetcode_941

import "testing"
import "reflect"

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[3,5,5], false},
		{[]int{1, 3, 5, 2, 1}, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := validMountainArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
