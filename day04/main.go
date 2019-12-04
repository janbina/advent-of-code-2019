package main

import "fmt"

type code []int

func main() {
	part12()
}

func getInput() (code, code) {
	return []int{1, 9, 7, 4, 8, 7}, []int{6, 7, 3, 2, 5, 1}
}

func (code code) inc() {
	for i := len(code) - 1; i >= 0; i-- {
		code[i] += 1
		if code[i] == 10 {
			code[i] = 0
		} else {
			break
		}
	}
}

func (code code) equals(code2 code) bool {
	if len(code) != len(code2) {
		return false
	}
	for i := 0; i < len(code); i++ {
		if code[i] != code2[i] {
			return false
		}
	}
	return true
}

func (code code) isValidP1() bool {
	prev := -1
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

func (code code) isValidP2() bool {
	prev := -1
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
	for !min.equals(max) {
		if min.isValidP1() {
			p1Count++
		}
		if min.isValidP2() {
			p2Count++
		}
		min.inc()
	}
	fmt.Println(p1Count)
	fmt.Println(p2Count)
}
