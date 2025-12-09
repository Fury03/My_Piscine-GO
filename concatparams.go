package student

func ConcatParams(args []string) string {
	result := ""

	for i, val := range args {
		result += val
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
