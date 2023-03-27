package stringsamples

import "strings"

// Сравнивает две строки и возвращает:
// '0' - если строки равны
// '1' - если первая строка раньше второй
// '-1' - если вторая строка раньше первой
func Compare(first string, second string) int {
	return strings.Compare(first, second)
}
