package student

func SplitWhiteSpaces(str string) []string {
	result := []string{}
	word := ""

	for _, char := range str {
		if char != ' ' && char != '\t' && char != '\n' {
			word += string(char)
		} else {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
