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
	start := 1
	var v1, v2 int
	var idx int
	for {
		op, pm1, pm2, _ := instruction(code[idx])
		switch op {
		case 1:
			if pm1 == 1 {
				v1 = code[idx+1]
			} else {
				v1 = code[code[idx+1]]
			}

			if pm2 == 1 {
				v2 = code[idx+2]
			} else {
				v2 = code[code[idx+2]]
			}

			code[code[idx+3]] = v1 + v2
			idx += 4
		case 2:
			if pm1 == 1 {
				v1 = code[idx+1]
			} else {
				v1 = code[code[idx+1]]
			}

			if pm2 == 1 {
				v2 = code[idx+2]
			} else {
				v2 = code[code[idx+2]]
			}

			code[code[idx+3]] = v1 * v2
			idx += 4
		case 3:
			code[code[idx+1]] = start
			idx += 2
		case 4:
			fmt.Println(code[code[idx+1]])
			idx += 2
		case 99:
			os.Exit(1)
		default:
			fmt.Printf("got invalid code: %d\n", op)
			os.Exit(1)
		}
	}
}

func instruction(c int) (op, pmode1, pmode2, pmode3 int) {
	code := fmt.Sprintf("%05d", c)

	op, _ = strconv.Atoi(string(code[3:5]))
	pmode1, _ = strconv.Atoi(string(code[2]))
	pmode2, _ = strconv.Atoi(string(code[1]))
	pmode3, _ = strconv.Atoi(string(code[0]))

	return op, pmode1, pmode2, pmode3
}

func readinput() []int {
	f, _ := os.Open("../input")
	// f, _ := os.Open("../input_t")
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
