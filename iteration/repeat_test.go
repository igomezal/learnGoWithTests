package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestOriginalRepeat(t *testing.T) {
	repeated := OriginalRepeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkOriginalRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OriginalRepeat("a", 5)
	}
}
