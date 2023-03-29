package stringsamples

import "strings"

// Проверяет содержит ли строка 'text' подстроку 'substring'
func Contains(text, substring string) bool {
	return strings.Contains(text, substring)
}
