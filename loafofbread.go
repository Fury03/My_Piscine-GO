package student

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	returnStr := ""
	count := 0

	for _, ch := range str {
		returnStr += string(ch)

		if ch != ' ' {
			count++
		}

		if count == 5 {
			returnStr += " "
			count = 0
		}
	}

	return returnStr + "\n"
}
