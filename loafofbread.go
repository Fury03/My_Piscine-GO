package student

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	returnStr := ""
	wordCount := 0

	for _, ch := range str {
		if ch == ' ' {
			continue 
		}

		returnStr += string(ch)
		wordCount++

		if wordCount == 5 {
			returnStr += " "
			wordCount = 0
		}
	}

	return returnStr + "\n"
}
