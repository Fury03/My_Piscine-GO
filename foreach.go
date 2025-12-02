package main

func ForEach(f func(int), a []int) {
	for _, ch := range a {
		f(ch)
	}
}
