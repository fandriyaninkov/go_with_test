package stringsamples

import "testing"

func TestContains(t *testing.T) {
	t.Run("search 'numer' in 'enumerating'", func(t *testing.T) {
		res := Contains("enumerating", "numer")
		expected := true
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("searc 'Dep' in 'Deeplomant'", func(t *testing.T) {
		res := Contains("Deeplomant", "Dep")
		expected := false
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("search empty string in 'hello'", func(t *testing.T) {
		res := Contains("hello", "")
		expected := true
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("search empty string in empty string", func(t *testing.T) {
		res := Contains("", "")
		expected := true
		assertBoolCorrectMessage(t, res, expected)
	})
}
