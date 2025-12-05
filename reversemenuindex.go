package student

func ReverseMenuIndex(menu []string) []string {
	l := len(menu)
	reversed := make([]string, l)

	for i := 0; i < l; i++ {
		reversed[i] = menu[l-1-i]
	}
	return reverse
}
