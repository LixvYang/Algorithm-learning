package leetcode

import "testing"
import "reflect"

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// res := SpiralOrder(matrix)
	// t.Log(res)

	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{matrix,[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SpiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}