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
