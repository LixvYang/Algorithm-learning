package leetcode

import (
	"testing"
)

func Test_Rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	matrix := [][]int{[[1,2,3],[4,5,6],[7,8,9]]}
	re := [][]int{[[7,4,1],[8,5,2],[9,6,3]]}
	tests := []struct {
		args args
		want [][]int
	}{
		{matrix, re}
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Rotate(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}