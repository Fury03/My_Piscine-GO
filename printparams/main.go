package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, param := range args {
		for _, ch := range param {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
