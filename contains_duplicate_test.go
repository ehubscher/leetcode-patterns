package main_test

import (
	"testing"

	lp "github.com/ehubscher/leetcode-patterns"
)

var numstests = []struct {
	in  []int
	out bool
}{
	{[]int{1,2,3,1}, true},
	{[]int{1,2,3,4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func TestContainsDuplicate(t *testing.T) {
	for _, tt := range numstests {
		var res bool = lp.ContainsDuplicate(tt.in)

		if res != tt.out {
			t.Errorf("got: %t, want: %t.", res, tt.out)
		}
	}
}
