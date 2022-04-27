package leetcode

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct{
		heights []int
		want int
	}{
		{[]int{2,1,5,6,2,3}, 10},
	}

	for _, tt := range tests {
		if got := LargestRectangleArea(tt.heights); got != tt.want {
			t.Errorf("LargestRectangleArea(%v) = %v, want %v", tt.heights, got, tt.want)
		}
	}
} 

// Benchmark .
func BenchmarkLargestRectangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LargestRectangleArea([]int{2,1,5,6,2,3})
	}
}