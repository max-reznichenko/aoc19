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

	res := 0
	for _, perm := range permutation([]int{0, 1, 2, 3, 4}) {
		am := amp(code, perm)
		if am > res {
			res = am
		}
	}

	fmt.Println(res)
}

func amp(code []int, perm []int) int {
	res := 0
	for _, b := range perm {
		res = compute(code, b, res)
	}

	return res
}

func compute(code []int, phase, input int) int {
	var idx int
	iseq := 0
	for {
		op, v1, v2, v3 := parseIns(code, idx)
		switch op {
		case 1:
			code[v3] = v1 + v2
			idx += 4
		case 2:
			code[v3] = v1 * v2
			idx += 4
		case 3:
			if iseq == 0 {
				code[v1] = phase
				iseq++
			} else {
				code[v1] = input
			}
			idx += 2
		case 4:
			return v1
			// idx += 2
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
			os.Exit(1) // this should never happen
			// return 0
		default:
			fmt.Printf("got invalid code: %d\n", op)
			os.Exit(1)
		}
	}
}

func parseIns(code []int, idx int) (op, v1, v2, v3 int) {
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

// copied from https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}
