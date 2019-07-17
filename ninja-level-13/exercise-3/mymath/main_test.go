package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		input    []int
		expected float64
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{1, 4, 6, 8, 100}, 6},
		test{[]int{0, 8, 10, 1000}, 9},
		test{[]int{9000, 4, 10, 8, 6, 12}, 9},
		test{[]int{123, 744, 140, 200}, 170},
	}

	for _, test := range tests {
		got := CenteredAvg(test.input)

		if got != test.expected {
			t.Error("Expected", test.expected, "Got", got)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8})
	}
}
