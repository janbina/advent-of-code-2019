package main

import "fmt"

func main() {
	part12()
}

func getInput() (int, int) {
	return 197487, 673251
}

func isValidP1(code string) bool {
	flag := false
	for i := 1; i < len(code); i++ {
		if code[i-1] > code[i] {
			return false
		}
		flag = flag || code[i-1] == code[i]
	}
	return flag
}

func isValidP2(code string) bool {
	flag := false
	same := 1
	for i := 1; i < len(code); i++ {
		if code[i-1] > code[i] {
			return false
		} else if code[i-1] == code[i] {
			same++
		} else {
			flag = flag || same == 2
			same = 1
		}
	}
	return flag || same == 2
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
