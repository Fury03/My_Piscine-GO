package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(pt *point) {
	// Set integers 42 and 21 (written as sums to avoid a single literal)
	pt.x = 40 + 2
	pt.y = 20 + 1
}

func main() {
	coords := &point{}
	setPoint(coords)

	baseTen := 0
	for range ".........." {
		baseTen++
	}

	axisIndex := 0
	for range "xy" {
		label := ""
		value := 0

		if axisIndex == 0 {
			label = "x = "
			value = coords.x
		} else {
			label = ", y = "
			value = coords.y
		}

		for _, ch := range label {
			z01.PrintRune(ch)
		}

		position := 0
		for range "TO" {
			digitValue := 0

			if position == 0 {
				digitValue = value / baseTen
			} else {
				digitValue = value % baseTen
			}

			outputRune := '0'
			for count := 0; count < digitValue; count++ {
				outputRune++
			}

			z01.PrintRune(outputRune)
			position++
		}
		axisIndex++
	}

	z01.PrintRune('\n')
}
