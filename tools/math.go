package tools

func Min(a int, b ...int) int {
	min := a
	for i := 0; i < len(b); i++ {
		if min > b[i] {
			min = b[i]
		}
	}
	return min
}

func Max(a int, b ...int) int {
	max := a
	for i := 0; i < len(b); i++ {
		if max < b[i] {
			max = b[i]
		}
	}
	return max
}
