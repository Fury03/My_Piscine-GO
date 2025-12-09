package student

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	result := []int{}
	for _, ch := range str {
		result = append(result, int(ch))
	}
	return result
}
