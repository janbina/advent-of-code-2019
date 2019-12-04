package main

import "fmt"

func main() {
	part12()
}

func getInput() (int, int) {
	return 197487, 673251
}

func isValidP1(code string) bool {
	var prev byte = 0
	flag := false
	for i := 0; i < len(code); i++ {
		if prev > code[i] {
			return false
		} else if prev == code[i] {
			flag = true
		}
		prev = code[i]
	}
	return flag
}

func isValidP2(code string) bool {
	var prev byte = 0
	flag := false
	same := 1
	for i := 0; i < len(code); i++ {
		if prev > code[i] {
			return false
		} else if prev == code[i] {
			same++
		} else {
			if same == 2 {
				flag = true
			}
			same = 1
		}
		prev = code[i]
	}
	if same == 2 {
		flag = true
	}
	return flag
}

func part12() {
	min, max := getInput()

	p1Count := 0
	p2Count := 0

	for codeInt := min; codeInt <= max; codeInt++ {
		code := fmt.Sprint(codeInt)
		if isValidP1(code) {
			p1Count++
		}
		if isValidP2(code) {
			p2Count++
		}
	}

	fmt.Println(p1Count)
	fmt.Println(p2Count)
}
