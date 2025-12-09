package student

func ToLower(s string) string {
	runes := []rune(s)
	for i, ch := range runes {
		if ch >= 'A' && ch <= 'Z' {
			runes[i] = ch + ('a' - 'A')
		}
	}
	return string(runes)
}
