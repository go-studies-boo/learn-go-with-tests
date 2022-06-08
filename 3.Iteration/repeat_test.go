package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 7)
	fmt.Println(repeated)
	// Output: bbbbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// NOTE by default Benchmarks are run sequentially.