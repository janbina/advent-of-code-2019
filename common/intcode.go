package common

func getOpCodeAndModes(i int) (int, [3]int) {
	return i % 100, [3]int{(i / 100) % 10, (i / 1000) % 10, (i / 10000) % 10}
}

func getOpAddr(mem map[int64]int64, modes [3]int, ip int64, offset int64, relativeBase int64) int64 {
	switch modes[offset] {
	case 0:
		return mem[ip+offset+1]
	case 1:
		return ip + offset + 1
	case 2:
		return relativeBase + mem[ip+offset+1]
	default:
		panic("Invalid mode")
	}
}

// RunIntcode program with memory *mem* and input provided by *in* channel
// Output will be written to *out* channel and termination signalized to *done* channel
func RunIntcode(mem map[int64]int64, in chan int64, out chan int64, done chan struct{}) {
	RunIntcodeInRequest(mem, in, nil, out, done)
}

func RunIntcodeInRequest(mem map[int64]int64, in chan int64, inRequest chan struct{}, out chan int64, done chan struct{}) {
	var ip int64 = 0
	var relativeBase int64 = 0

	for {
		opCode, modes := getOpCodeAndModes(int(mem[ip]))
		op0addr := getOpAddr(mem, modes, ip, 0, relativeBase)
		op1addr := getOpAddr(mem, modes, ip, 1, relativeBase)
		op2addr := getOpAddr(mem, modes, ip, 2, relativeBase)

		switch opCode {
		case 1:
			mem[op2addr] = mem[op0addr] + mem[op1addr]
			ip += 4
		case 2:
			mem[op2addr] = mem[op0addr] * mem[op1addr]
			ip += 4
		case 3:
			if inRequest != nil {
				inRequest <- struct{}{}
			}
			mem[op0addr] = <-in
			ip += 2
		case 4:
			if out != nil {
				out <- mem[op0addr]
			}
			ip += 2
		case 5:
			if mem[op0addr] != 0 {
				ip = mem[op1addr]
			} else {
				ip += 3
			}
		case 6:
			if mem[op0addr] == 0 {
				ip = mem[op1addr]
			} else {
				ip += 3
			}
		case 7:
			if mem[op0addr] < mem[op1addr] {
				mem[op2addr] = 1
			} else {
				mem[op2addr] = 0
			}
			ip += 4
		case 8:
			if mem[op0addr] == mem[op1addr] {
				mem[op2addr] = 1
			} else {
				mem[op2addr] = 0
			}
			ip += 4
		case 9:
			relativeBase += mem[op0addr]
			ip += 2
		case 99:
			if out != nil {
				close(out)
			}
			done <- struct{}{}
			return
		default:
			panic("Invalid command")
		}
	}
}

// RunIntcodeSimple is wrapper over RunIntcode without channels
// Input provided as a slice, output returned as a slice
func RunIntcodeSimple(mem map[int64]int64, input []int64) []int64 {
	var output []int64

	in := make(chan int64)
	inRequest := make(chan struct{})
	out := make(chan int64)
	done := make(chan struct{})

	go RunIntcodeInRequest(mem, in, inRequest, out, done)

	inCnt := 0

	for {
		select {
		case <-inRequest:
			in <- input[inCnt]
			inCnt++
		case x := <-out:
			output = append(output, x)
		case <-done:
			return output
		}
	}
}
