package student

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	sortNum(arr)
	avg := (len(arr) + 1) / 2
	return arr[avg-1]
}

func sortNum(num []int) {
	length := len(num)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
}
