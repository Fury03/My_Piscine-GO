package student

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, char := range a {
		for _, rune := range char {
			z01.PrintRune(rune)
		}
		z01.PrintRune('\n')
	}
}
