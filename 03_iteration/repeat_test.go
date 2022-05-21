package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := strings.Repeat("a", 5)

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
