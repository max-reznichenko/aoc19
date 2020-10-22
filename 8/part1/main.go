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

	inX, inY := 25, 6

	layers := [][]int{}
	layer := []int{}
	for i, d := range input {
		if i%(inX*inY) == 0 && i != 0 {
			layers = append(layers, layer)
			layer = []int{}
		}

		layer = append(layer, d)

		if i+1 == len(input) {
			layers = append(layers, layer)
		}
	}

	var zeroCounter int
	var res int
	var m map[int]int

	for _, layer := range layers {
		m = map[int]int{}
		for _, d := range layer {
			m[d]++
		}
		if zeroCounter == 0 || m[0] < zeroCounter {
			zeroCounter = m[0]
			res = m[1] * m[2]
		}
	}

	fmt.Println(res)
}

func readinput() []int {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []int{}

	scanner.Scan()
	inputstr := scanner.Text()

	for _, s := range strings.Split(inputstr, "") {
		l, _ := strconv.Atoi(s)
		input = append(input, l)
	}
	return input
}
