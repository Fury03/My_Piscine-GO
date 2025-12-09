package student

func ToUpper(s string) string {
	result := []rune(s)

	for i, r := range result {
		// Check if r is a lowercase letter
		if r >= 'a' && r <= 'z' {
			result[i] = r - 32 // convert to uppercase using ASCII
		}
	}

	return string(result)
}
