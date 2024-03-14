package main_test

import (
	"testing"

	lp "github.com/ehubscher/leetcode-patterns"
)

var missNums = []struct {
	in  []int
	out int
}{
	{[]int{3,1,0}, 2},
	{[]int{0,1}, 2},
	{[]int{9,6,4,2,3,5,7,0,1}, 8},
}

func TestMissingNumber(t *testing.T) {
	for _, tt := range missNums {
		var res int = lp.MissingNumber(tt.in)

		if res != tt.out {
			t.Errorf("got: %d, want: %d.", res, tt.out)
		}
	}
}
