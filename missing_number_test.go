package main_test

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
	"time"

	lp "github.com/ehubscher/leetcode-patterns"
)

var missNums = []struct {
	in  []int
	out int
}{
	{nil, 0},
	{[]int{}, 0},
	{[]int{0}, 1},
	{[]int{1}, 0},
	{[]int{2}, -1},
	{[]int{0,1}, 2},
	{[]int{0,1,2}, 3},
	{[]int{3,1,0}, 2},
	{[]int{9,6,4,2,3,5,7,0,1}, 8},
}

func TestMissingNumber(t *testing.T) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 3, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "Function Name\tDuration\tInput\tOutput")

	for _, tt := range missNums {
		inCpy := append([]int{}, tt.in...)

		start := time.Now()
		var res int = lp.MissingNumber(inCpy)
		duration := time.Since(start)

		fmt.Fprintf(writer, "%s\t%v\t%v\t%d\n", "MissingNumber()", duration, tt.in, tt.out)

		if res != tt.out {
			t.Errorf("got: %d, want: %d.\n", res, tt.out)
		}
	}

	fmt.Fprintln(writer, "")
	writer.Flush()
}
