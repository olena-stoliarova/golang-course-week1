package pangram

import "strings"

// IsPangram checks if given string is a pangram
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for _, char := range "abcdefghijklmnopqrstuvwxyz" {
		if !strings.Contains(s, string(char)) {
			return false
		}
	}
	// Тут должно быть решение
	// написав код - необходимо запустить тесты
	// Эти комментарии можно удалять
	return true
}
