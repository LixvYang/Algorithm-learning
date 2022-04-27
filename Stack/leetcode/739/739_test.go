package leetcode

import "testing"

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		want       []int
	}{
		{[]int{73,74,75,71,69,72,76,73}, []int{1,1,4,2,1,1,0,0}},
	}

	for _, tt := range tests {
		if got := DailyTemperatures(tt.temperatures); !eql(got, tt.want) {
				t.Fatalf("got: %v, want: %v", got, tt.want)
			}
	}
}

func BenchmarkDailyTemperatures(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DailyTemperatures([]int{73,74,75,71,69,72,76,73})
	}
}

func eql(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}