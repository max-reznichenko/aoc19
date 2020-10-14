package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	w1, w2 := readinput()
	var pm [2][][3]int // [x, y, v]

	for wi, w := range [2][]string{w1, w2} {
		x, y := 0, 0
		for _, d := range w {
			s, _ := strconv.Atoi(d[1:])
			for i := 0; i < s; i++ {
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

				pm[wi] = append(pm[wi], [3]int{x, y, 1})
			}
		}
	}

	min := 0
	for _, p1 := range pm[0] {
		for _, p2 := range pm[1] {
			if p2 == p1 {
				d := md([]int{0, 0}, []int{p1[0], p1[1]})
				if min == 0 {
					min = d
				}
				if d < min {
					min = d
				}
			}
		}
	}

	fmt.Println(min)
}

func md(p1, p2 []int) int {
	d1 := int(math.Abs(float64(p1[0]) - float64(p2[0])))
	d2 := int(math.Abs(float64(p1[1]) - float64(p2[1])))
	return d1 + d2
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
