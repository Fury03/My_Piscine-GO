package student

func FindNextPrime(nb int) int {
	if IsCurrentPrime(nb) { //recursion to find prime
		return nb
	}
	return FindNextPrime(nb + 1)
}

func IsCurrentPrime(nb int) bool { //recalling the isprime func from the last exercise just writing it all out to work here
	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
