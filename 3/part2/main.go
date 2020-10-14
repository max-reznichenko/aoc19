package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	w1, w2 := readinput()
	var pm [2][][3]int // [x, y, step]

	for wi, w := range [2][]string{w1, w2} {
		x, y, step := 0, 0, 0
		for _, d := range w {
			s, _ := strconv.Atoi(d[1:])
			for i := 0; i < s; i++ {
				step++
				switch string(d[0]) {
				case "R":
					x++
				case "L":
					x--
				case "U":
					y++
				case "D":
					y--
				}
				pm[wi] = append(pm[wi], [3]int{x, y, step})
			}
		}
	}

	shortestPath := 0
	for _, point1 := range pm[0] {
		for _, point2 := range pm[1] {
			if point1[0] == point2[0] && point1[1] == point2[1] {
				stepsSum := point1[2] + point2[2]
				if shortestPath == 0 {
					shortestPath = stepsSum
				}
				if stepsSum < shortestPath {
					shortestPath = stepsSum
				}
			}
		}
	}

	fmt.Println(shortestPath)
}

func readinput() ([]string, []string) {
	f, _ := os.Open("../input")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	w1, w2 := []string{}, []string{}

	scanner.Scan()
	inputstr := scanner.Text()
	for _, s := range strings.Split(inputstr, ",") {
		w1 = append(w1, s)
	}

	scanner.Scan()
	inputstr = scanner.Text()
	for _, s := range strings.Split(inputstr, ",") {
		w2 = append(w2, s)
	}
	return w1, w2
}
