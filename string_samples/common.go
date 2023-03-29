package stringsamples

import "testing"

func assertBoolCorrectMessage(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}
