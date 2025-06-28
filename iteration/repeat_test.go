package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q, want %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	//... setup ...
	for b.Loop() {
		//... code to measure ...
		Repeat("a")
	}
	//... cleanup ...
}
