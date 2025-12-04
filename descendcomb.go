package student

import "github.com/01-edu/z01"

func PrintNum(a int) {
	first := a / 10
	second := a % 10

	z01.PrintRune(rune(first) + '0')
	z01.PrintRune(rune(second) + '0')
}

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {

			PrintNum(i)
			z01.PrintRune(' ')
			PrintNum(j)

			if i == 1 && j == 0 {
				return
			}

			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
