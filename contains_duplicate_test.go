package main_test

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
	"time"

	lp "github.com/ehubscher/leetcode-patterns"
)

var dupNums = []struct {
	in  []int
	out bool
}{
	{[]int{1,2,3,1}, true},
	{[]int{1,2,3,4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func TestContainsDuplicate(t *testing.T) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 3, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "Function Name\tDuration\tInput\tOutput")
	
	for _, tt := range dupNums {
		inCpy := append([]int{}, tt.in...)

		start := time.Now()
		var res bool = lp.ContainsDuplicate(inCpy)
		duration := time.Since(start)

		fmt.Fprintf(writer, "%s\t%v\t%v\t%v\n", "ContainsDuplicate()", duration, tt.in, tt.out)

		if res != tt.out {
			t.Errorf("got: %t, want: %t.\n", res, tt.out)
		}
	}

	fmt.Fprintln(writer, "")
	writer.Flush()
}
