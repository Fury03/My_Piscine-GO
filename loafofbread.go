package student

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	// Remove spaces for counting letters
	letters := ""
	for _, ch := range str {
		if ch != ' ' {
			letters += string(ch)
		}
	}

	returnStr := ""
	for i, ch := range letters {
		returnStr += string(ch)
		// Add a space after every 5 letters
		if (i+1)%5 == 0 && i != len(letters)-1 {
			returnStr += " "
		}
	}

	return returnStr + "\n"
}
