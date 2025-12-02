package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a1, ok1 := atoi(os.Args[1])
	op := os.Args[2]
	a2, ok2 := atoi(os.Args[3])
	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		printInt(a1 + a2)
	case "-":
		printInt(a1 - a2)
	case "*":
		printInt(a1 * a2)
	case "/":
		if a2 == 0 {
			writeString("No division by 0\n")
		} else {
			printInt(a1 / a2)
		}
	case "%":
		if a2 == 0 {
			writeString("No modulo by 0\n")
		} else {
			printInt(a1 % a2)
		}
	}
}

// Manual string to int conversion
func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	}
	n := 0
	for i := start; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n * sign, true
}

// Print int to stdout
func printInt(n int) {
	if n == 0 {
		writeByte('0')
		writeByte('\n')
		return
	}
	if n < 0 {
		writeByte('-')
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append(digits, byte(n%10)+'0')
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		writeByte(digits[i])
	}
	writeByte('\n')
}

// Helper to write string to stdout
func writeString(s string) {
	for i := 0; i < len(s); i++ {
		writeByte(s[i])
	}
}

// Helper to write single byte to stdout
func writeByte(b byte) {
	os.Stdout.Write([]byte{b})
}
