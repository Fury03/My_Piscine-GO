package student

func Compact(ptr *[]string) int {
	closed := []string{}
	for _, v := range *ptr {
		if v != "" {
			closed = append(closed, v)
		}
	}
	*ptr = closed
	return len(closed)
}
