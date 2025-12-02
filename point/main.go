package main

import (

	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

const (
	LetX = 'x'
	LetY = 'y'
	EqualTo = '='
	comma = ','
	digit4 = '4'
	digit2 = '2'
	digit1 = '1'
	newline = '\n'
	space = ' '

)

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	p := &point{}
	setPoint(p)
	output := []rune{
		LetX, space, EqualTo, space, digit4, digit2,
		 comma, space, LetY, space, EqualTo, digit2,digit1,
		 newline,
	}
	for _, ch := range output {
	z01.PrintRune(ch)
	}

}
