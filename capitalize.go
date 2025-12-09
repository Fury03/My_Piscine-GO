package student

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true // shows the start of a new word

	for i, ch := range runes {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			if newWord {
				// capitalize first letter if it's a lowercase letter
				if ch >= 'a' && ch <= 'z' {
					runes[i] = ch - ('a' - 'A')
				}
				newWord = false
			} else {
				// lowercase the rest of letters if uppercase
				if ch >= 'A' && ch <= 'Z' {
					runes[i] = ch + ('a' - 'A')
				}
			}
		} else {
			// non-alphanumeric character → next char starts a new word
			newWord = true
		}
	}

	return string(runes)
}
