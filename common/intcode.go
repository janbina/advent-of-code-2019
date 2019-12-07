package common

func getOpCodeAndModes(i int) (int, [3]int) {
	return i % 100, [3]int{(i / 100) % 10, (i / 1000) % 10, (i / 10000) % 10}
}

func getVal(mem []int, modes [3]int, ip int, offset int) int {
	if modes[offset] == 1 {
		return mem[ip+offset+1]
	} else {
		return mem[mem[ip+offset+1]]
	}
}

func RunIntcode(mem []int, ip int, input []int) (int, []int) {
	inCnt := 0
	var output []int

	for {
		opCode, modes := getOpCodeAndModes(mem[ip])

		switch opCode {
		case 1:
			mem[mem[ip+3]] = getVal(mem, modes, ip, 0) + getVal(mem, modes, ip, 1)
			ip += 4
		case 2:
			mem[mem[ip+3]] = getVal(mem, modes, ip, 0) * getVal(mem, modes, ip, 1)
			ip += 4
		case 3:
			if inCnt >= len(input) {
				return ip, output
			}
			mem[mem[ip+1]] = input[inCnt]
			inCnt++
			ip += 2
		case 4:
			output = append(output, getVal(mem, modes, ip, 0))
			ip += 2
		case 5:
			if getVal(mem, modes, ip, 0) != 0 {
				ip = getVal(mem, modes, ip, 1)
			} else {
				ip += 3
			}
		case 6:
			if getVal(mem, modes, ip, 0) == 0 {
				ip = getVal(mem, modes, ip, 1)
			} else {
				ip += 3
			}
		case 7:
			if getVal(mem, modes, ip, 0) < getVal(mem, modes, ip, 1) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
			ip += 4
		case 8:
			if getVal(mem, modes, ip, 0) == getVal(mem, modes, ip, 1) {
				mem[mem[ip+3]] = 1
			} else {
				mem[mem[ip+3]] = 0
			}
			ip += 4
		case 99:
			return ip, output
		default:
			panic("Invalid command")
		}
	}
}
