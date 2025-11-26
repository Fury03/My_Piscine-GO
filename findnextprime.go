package student

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	// If nb is even, make it odd
	if nb%2 == 0 {
		nb++
	}

	for {
		if isPrime(nb) {
			return nb
		}
		nb += 2 // only check odd numbers
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Check up to sqrt(n)
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
