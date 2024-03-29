package util

func MapSlice[T, U any](slice []T, f func(T) U) []U {
	mapped := make([]U, len(slice))
	for i := range slice {
		mapped[i] = f(slice[i])
	}
	return mapped
}

func MapSliceWithError[T, U any](slice []T, f func(T) (U, error)) ([]U, error) {
	mapped := make([]U, len(slice))
	for i := range slice {
		res, err := f(slice[i])
		if err != nil {
			return nil, err
		}
		mapped[i] = res
	}
	return mapped, nil
}

func ReduceSlice[T, U any](slice []T, f func(T, U) U) U {
	var accumulator U
	for i := range slice {
		accumulator = f(slice[i], accumulator)
	}
	return accumulator
}

func SliceEvery[T any](slice []T, f func(T) bool) bool {
	for i := range slice {
		if !f(slice[i]) {
			return false
		}
	}
	return true
}

func SliceFilter[T any](s []T, f func(T) bool) []T {
	var filtered []T
	for _, el := range s {
		if f(el) {
			filtered = append(filtered, el)
		}
	}
	return filtered
}

func SumIntSlice(slice []int) (total int) {
	for _, v := range slice {
		total += v
	}
	return
}

func MaxIntSlice(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
