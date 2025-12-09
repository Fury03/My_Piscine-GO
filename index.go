package student

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}

	runesS := []rune(s)
	runesFind := []rune(toFind)
	lenS := len(runesS)
	lenFind := len(runesFind)

	for i := 0; i <= lenS-lenFind; i++ {
		match := true
		for j := 0; j < lenFind; j++ {
			if runesS[i+j] != runesFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
