package stringsamples

import "testing"

func TestContainsAny(t *testing.T) {
	t.Run("search any of 'i' in 'team'", func(t *testing.T) {
		res := ContainsAny("team", "i")
		expected := false
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("search any of 'ui' in 'fail'", func(t *testing.T) {
		res := ContainsAny("fail", "ui")
		expected := true
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("search any of 'ui' in 'ure'", func(t *testing.T) {
		res := ContainsAny("ure", "ui")
		expected := true
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("search any of 'ui' in 'failure'", func(t *testing.T) {
		res := ContainsAny("failure", "ui")
		expected := true
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("search empty string in 'foo'", func(t *testing.T) {
		res := ContainsAny("foo", "")
		expected := false
		assertBoolCorrectMessage(t, res, expected)
	})
	t.Run("search empty string in empty string", func(t *testing.T) {
		res := ContainsAny("", "")
		expected := false
		assertBoolCorrectMessage(t, res, expected)
	})
}
