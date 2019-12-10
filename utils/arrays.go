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

func StringsToInts64(arr []string) []int64 {
	r := make([]int64, len(arr))
	for i, s := range arr {
		r[i] = ToInt64(s)
	}
	return r
}

func IntSliceToMap(arr []int64) map[int64]int64 {
	mem := make(map[int64]int64)
	for i, v := range arr {
		mem[int64(i)] = v
	}
	return mem
}
