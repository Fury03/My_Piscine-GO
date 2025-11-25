package student

func LastRune(s string) rune {
	char := []rune(s)
	return char[len(char)-1]
}
