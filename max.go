package student

func Max(a []int) int {
	highestval := 0
	for i := 0; i < len(a); i++ {
		if a[i] > highestval {
			highestval = a[i]
		}
	}
	return highestval
}
