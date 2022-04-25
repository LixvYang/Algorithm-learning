package leetcode

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for _, tt := range tests {
		if got := Trap(tt.nums); got != tt.want {
			t.Fatalf("trap(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

func TestTrap2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for _, tt := range tests {
		if got := Trap2(tt.nums); got != tt.want {
			t.Fatalf("trap(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}

// Benchmark .
func BenchmarkTrap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := Trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		b.Logf("sum = %v\n", sum)
	}
}

func BenchmarkTrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		b.Logf("sum = %v\n", sum)
	}
}