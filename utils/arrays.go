package utils

func CopyInts(arr []int) []int {
	r := make([]int, len(arr))
	copy(r, arr)
	return r
}

func StringsToInts(arr []string) []int {
	r := make([]int, len(arr))
	for i, s := range arr {
		r[i] = ToInt(s)
	}
	return r
}
