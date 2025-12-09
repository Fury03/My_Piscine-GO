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
		if res, ok := safeAdd(a1, a2); ok {
			printInt(res)
		}
	case "-":
		if res, ok := safeSub(a1, a2); ok {
			printInt(res)
		}
	case "*":
		if res, ok := safeMul(a1, a2); ok {
			printInt(res)
		}
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
		// Check for overflow while building the number
		if n > (9223372036854775807-int(c-'0'))/10 {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n * sign, true
}

func safeAdd(a, b int) (int, bool) {
	sum := a + b
	if (b > 0 && sum < a) || (b < 0 && sum > a) {
		return 0, false
	}
	return sum, true
}

// Safe subtraction with overflow check
func safeSub(a, b int) (int, bool) {
	diff := a - b
	if (b < 0 && diff < a) || (b > 0 && diff > a) {
		return 0, false
	}
	return diff, true
}

func safeMul(a, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	res := a * b
	if res/b != a {
		return 0, false
	}
	return res, true
}

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

func writeString(s string) {
	for i := 0; i < len(s); i++ {
		writeByte(s[i])
	}
}

func writeByte(b byte) {
	os.Stdout.Write([]byte{b})
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
		//} else {
		//	return false
	} //
}

func IsNegative(n int) bool {
	if n < 0 {
		return true
	}
	return false
}
