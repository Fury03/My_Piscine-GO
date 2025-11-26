package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	// Check for flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	// Process args and print letters/spaces
	for _, a := range args {
		n := atoi(a) // custom atoi (returns -1 on invalid)
		if n >= 1 && n <= 26 {
			var r rune
			if upper {
				r = rune('A' + n - 1)
			} else {
				r = rune('a' + n - 1)
			}
			z01.PrintRune(r)
		} else {
			z01.PrintRune(' ')
		}
	}

	// Print newline only if we actually processed at least one argument
	if len(args) > 0 {
		z01.PrintRune('\n')
	}
}

// Custom Atoi that returns -1 for invalid input
func atoi(s string) int {
	sign := 1
	start := 0
	result := 0

	if len(s) == 0 {
		return -1
	}

	// Handle optional + or -
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}
