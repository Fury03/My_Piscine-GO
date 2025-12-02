/*package student

func Map(f func(int) bool, a []int) []bool {
	var boolArray []bool
	for _, v := range a {
		boolArray = make(boolArray, f(v))
	}
	return boolArray
}
*/

package student

func Map(f func(int) bool, a []int) []bool {
	boolArray := make([]bool, len(a))

	for i, v := range a {
		boolArray[i] = f(v)
	}

	return boolArray
}
