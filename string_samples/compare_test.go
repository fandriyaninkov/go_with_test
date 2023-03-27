package stringsamples

import (
	"fmt"
	"testing"
)

func ExampleCompare() {
	res := Compare("a", "b")
	fmt.Println(res)
	// Output: -1
}

func CompareTest(t *testing.T) {
	res := Compare("a", "a")
	epected := 0

	if res != epected {
		t.Errorf("expected '%d' but got '%d'", epected, res)
	}
}

func CompareTest2(t *testing.T) {
	res := Compare("b", "a")
	expected := 1
	if res != expected {
		t.Errorf("expected '%d' but got '%d'", expected, res)
	}
}

func CompareTest3(t *testing.T) {
	res := Compare("adult", "advantage")
	expected := 1
	if res != expected {
		t.Errorf("expected '%d' but got '%d'", expected, res)
	}
}
