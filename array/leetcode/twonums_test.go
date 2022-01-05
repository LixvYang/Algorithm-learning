package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var q []int = []int{2, 7, 11, 15}
	var target = 9
	if TwoSum(q, target) != nil {
		t.Log("ok")
	}

}
