package go_generics

func MinMax[T int | float64](slice []T) (T, T) {
	if len(slice) == 0 {
		return 0, 0
	}

	min := slice[0]
	max := slice[0]

	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}
