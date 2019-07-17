package dog

import (
	"fmt"
	"testing"
)

type test struct {
	input    int
	expected int
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func TestYears(t *testing.T) {

	testYears := []test{
		test{7, 49},
		test{0, 0},
		test{1, 7},
		test{10, 70},
		test{100, 700},
		test{5, 35},
	}

	for _, year := range testYears {
		testYears := Years(year.input)
		if testYears != year.expected {
			t.Error("Expected", year.expected, "Got", testYears)
		}
	}
}
