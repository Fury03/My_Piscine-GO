package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(p *point) {
	p.x = 42
	p.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	p := &point{}
	setPoint(p)

	printStr("x = ")
	printNbr(p.x)
	printStr(", y = ")
	printNbr(p.y)
	z01.PrintRune('\n')
}
