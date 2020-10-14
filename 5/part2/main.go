package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	code := readinput()
	start := 5

	var idx int
	for {
		op, v1, v2, v3 := instruction(code, idx)
		switch op {
		case 1:
			code[v3] = v1 + v2
			idx += 4
		case 2:
			code[v3] = v1 * v2
			idx += 4
		case 3:
			code[v1] = start
			idx += 2
		case 4:
			fmt.Println(v1)
			idx += 2
		case 5:
			if v1 > 0 {
				idx = v2
			} else {
				idx += 3
			}

		case 6:
			if v1 == 0 {
				idx = v2
			} else {
				idx += 3
			}
		case 7:
			if v1 < v2 {
				code[v3] = 1
			} else {
				code[v3] = 0
			}
			idx += 4
		case 8:
			if v1 == v2 {
				code[v3] = 1
			} else {
				code[v3] = 0
			}
			idx += 4
		case 99:
			os.Exit(1)
		default:
			fmt.Printf("got invalid code: %d\n", op)
			os.Exit(1)
		}
	}
}

func instruction(code []int, idx int) (op, v1, v2, v3 int) {
	ins := fmt.Sprintf("%05d", code[idx])

	op, _ = strconv.Atoi(string(ins[3:]))
	pmode1, _ := strconv.Atoi(string(ins[2]))
	pmode2, _ := strconv.Atoi(string(ins[1]))
	_, _ = strconv.Atoi(string(ins[0])) // not used

	switch op {
	case 1, 2, 5, 6, 7, 8:
		if pmode1 == 1 {
			v1 = code[idx+1]
		} else {
			v1 = code[code[idx+1]]
		}

		if pmode2 == 1 {
			v2 = code[idx+2]
		} else {
			v2 = code[code[idx+2]]
		}
		return op, v1, v2, code[idx+3]
	case 3:
		return op, code[idx+1], 0, 0
	case 4:
		if pmode1 == 1 {
			v1 = code[idx+1]
		} else {
			v1 = code[code[idx+1]]
		}
		return op, v1, 0, 0
	case 99:
		return op, 0, 0, 0
	default:
		return 0, 0, 0, 0
	}
}

func readinput() []int {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []int{}

	scanner.Scan()
	inputstr := scanner.Text()

	for _, s := range strings.Split(inputstr, ",") {
		l, _ := strconv.Atoi(s)
		input = append(input, l)
	}
	return input
}
