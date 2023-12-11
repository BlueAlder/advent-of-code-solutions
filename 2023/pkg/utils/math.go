package util

func GCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LCM(a, b int, nums ...int) int {
	res := (a * b) / GCD(a, b)

	for _, num := range nums {
		res = LCM(res, num)
	}
	return res
}
