package utils

func Min(args ...int) int {
	min := args[0]
	for _, x := range args {
		if x < min {
			min = x
		}
	}
	return min
}

func Max(args ...int) int {
	max := args[0]
	for _, x := range args {
		if x > max {
			max = x
		}
	}
	return max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDist(x1 int, x2 int, y1 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func Gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}
