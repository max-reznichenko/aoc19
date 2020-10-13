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
		t += sum(0, l)
	}
	fmt.Println(t)
}

func sum(s, f int) int {
	f = f/3 - 2
	if f >= 0 {
		return sum(s+f, f)
	}
	return s
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
