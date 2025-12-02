package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	p := &point{}
	setPoint(p)

	out := "x = 42, y = 21\n"
	for _, ch := range out {
		z01.PrintRune(ch)
	}
}
