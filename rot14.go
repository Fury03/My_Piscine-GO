package student

func Rot14(s string) string {
	result := ""

	for _, char := range s {

		if char >= 'a' && char <= 'z' {
			char = char + 14
			if char > 'z' {
				char = char - 26
			}
		}

		if char >= 'A' && char <= 'Z' {
			char = char + 14
			if char > 'Z' {
				char = char - 26
			}
		}

		result += string(char)
	}

	return result
}
