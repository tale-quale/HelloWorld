package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("count is greater then 0", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"

		assertCorrectMsg(t, repeated, expected)
	})
	t.Run("count is lower or equal to 0", func(t *testing.T) {
		repeated := Repeat("a", -7)
		expected := ""

		assertCorrectMsg(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	//... setup ...
	for b.Loop() {
		//... code to measure ...
		Repeat("a", 5)
	}
	//... cleanup ...
}

func ExampleRepeat() {
	repeated := Repeat("CA", 3)
	fmt.Println(repeated)
	// Output: CACACA
}

func assertCorrectMsg(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("repeated %q, expected %q", repeated, expected)
	}
}
