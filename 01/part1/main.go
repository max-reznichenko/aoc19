package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t int
	for _, l := range readinput() {
		t += l/3 - 2
	}

	fmt.Println(t)
}

func readinput() []int {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []int{}

	for scanner.Scan() {
		l, _ := strconv.Atoi(scanner.Text())
		input = append(input, l)
	}
	return input
}
