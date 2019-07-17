package abs

import (
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-1)

	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(i * -1)
	}
}
