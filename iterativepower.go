package student

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	answer := 1
	for i := 0; i < power; i++ {
		answer *= nb
	}

	return answer
}
