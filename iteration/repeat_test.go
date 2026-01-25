package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", repeatCount)
	s := ""
	for i := 0; i < repeatCount; i++ {
		s += "a"
	}

	expected := s

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", repeatCount)
	}
}
