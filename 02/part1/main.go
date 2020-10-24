package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readinput()
	input[1] = 12
	input[2] = 2
	var idx int
	for {
		switch input[idx] {
		case 1:
			input[input[idx+3]] = input[input[idx+1]] + input[input[idx+2]]
			idx += 4
		case 2:
			input[input[idx+3]] = input[input[idx+1]] * input[input[idx+2]]
			idx += 4
		case 99:
			fmt.Println(input[0])
			os.Exit(1)
		}
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
