package utils

func CopyInts(arr []int) []int {
	r := make([]int, len(arr))
	copy(r, arr)
	return r
}
