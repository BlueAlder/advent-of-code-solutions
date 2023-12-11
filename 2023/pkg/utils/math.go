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

// Checks if the int x is between or equal to
// the values a and b
func EqualOrBetween(a, b, x int) bool {
	return a <= x && x <= b ||
		(a-x)*(b-x) <= 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
